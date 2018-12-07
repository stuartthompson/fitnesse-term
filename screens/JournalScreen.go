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

package screens

import (
	"github.com/stuartthompson/fitnesse-term/configuration"
	"github.com/stuartthompson/fitnesse-term/data"
	"github.com/stuartthompson/fitnesse-term/io/screen"
)

// JournalScreen ...
type JournalScreen struct {
	screen			*screen.Screen
	configuration	*configuration.AppConfig
	journal 		*data.Journal
}

// NewJournalScreen ...
// Instantiates a new journal screen.
func NewJournalScreen(configuration *configuration.AppConfig, viewport *screen.Viewport, journal *data.Journal) *JournalScreen {
	screenStyle := &screen.Style{ShowBorder: true, BorderColor: 100}
	screen := screen.NewScreen(viewport, screenStyle)

	return &JournalScreen{screen: screen, configuration: configuration, journal: journal}
}

// Render ...
// Renders the config screen.
func (s *JournalScreen) Render() {
	s.screen.Clear()

	s.screen.RenderText("Journal", 1, 1, 255, 0)
	s.screen.RenderText("Version: "+s.journal.Version, 1, 3, 255, 0)
}
