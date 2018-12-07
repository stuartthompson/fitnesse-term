// Copyright 2018 Stuart Thompson.

// This file is part of Fitnesse-Term.

// Fitnesse-Term is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.

// Fitnesse-Term is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with Fitnesse-Term. If not, see <http://www.gnu.org/licenses/>.

package data

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"path"
)

// journalFileName ...
// The name of the fitness journal file.
const journalFileName = "journal.json"

// Journal ...
// Represents the journal of fitness data recorded by the user.
type Journal struct {
	Version string       		`json:version`
	Entries []JournalEntry		`json:Entries`
}

// JournalEntry ...
// Represents an entry in the fitness journal.
type JournalEntry struct {
	Timestamp string			`json:timestamp`
	Name string 				`json:timestamp`
}

// ReadJournal ...
// Reads the fitness journal from disk.
func (j *Journal) ReadJournal() error {
	journalFilePath, err := j.buildJournalFilePath()
	if err != nil {
		log.Print("Error building journal file path.")
		return err
	}

	// Create new (empty) journal if it does not exist
	if _, err := os.Stat(journalFilePath); os.IsNotExist(err) {
		j.writeNewJournal(journalFilePath)
	}

	// Read journal
	rawJournal, err := ioutil.ReadFile(journalFilePath)
	if err != nil {
		log.Fatal(err)
		return err
	}

	// Unmarshal journal
	var journal *Journal
	json.Unmarshal(rawJournal, &j)

	return nil
}

// buildJournalFilePath ...
// Builds the file path for the fitness journal.
func (j *Journal) buildJournalFilePath() (string, error) {
	// Get user's home directory
	usr, err := user.Current()
	if err != nil {
		log.Print("Error getting current user. Error: ", err)
		return "", err
	}

	// Build journal file path
	journalFilePath := path.Join(usr.HomeDir, journalFileName)

	return journalFilePath, nil
}

// writeNewJournal ...
// Writes a new (empty) journal file.
func (j *Journal) writeNewJournal(journalFilePath string) {
	log.Print("Writing new journal")
	journal := Journal{Version: "1.0"}
	journalJSON, err := json.Marshal(journal)
	if err != nil {
		log.Print("Error marshaling json.")
		log.Fatal(err)
	}
	err = ioutil.WriteFile(journalFilePath, journalJSON, 0666)
	if err != nil {
		log.Print("Error writing new (empty) journal.")
		log.Fatal("Error is: ", err)
	}
}
