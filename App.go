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

package main

import (
	"log"

	termbox "github.com/nsf/termbox-go"
	"github.com/stuartthompson/fitnesse-term/configuration"
	"github.com/stuartthompson/fitnesse-term/data"
	"github.com/stuartthompson/fitnesse-term/io"
	"github.com/stuartthompson/fitnesse-term/io/screen"
	"github.com/stuartthompson/fitnesse-term/screens"
)

// Screen ...
// Typedef for screen types.
type Screen int

// Defines screen types.
const (
	ConfigScreen = iota
	AboutScreen
	JournalScreen
)

// configFileName ...
// The name of the application configuration file.
const configFileName = ".fitnesserc"

// App ...
// Encapsulates main application logic.
type App struct {
	isRunning       bool
	eventListener   *io.EventListener
	configuration   *configuration.AppConfig
	journal			*data.Journal
	currentScreen   Screen
	configScreen    *screens.ConfigScreen
	aboutScreen     *screens.AboutScreen
	journalScreen	*screens.JournalScreen
	bottomBar       *screens.BottomBarComponent
}

// NewApp ...
// Initializes a new application instance.
func NewApp() *App {
	app := &App{
		isRunning:     true,
		configuration: &configuration.AppConfig{},
		journal: &data.Journal{},
		currentScreen: AboutScreen,
	}
	app.eventListener = io.NewEventListener(app.Render)

	return app
}

// Run ...
// Runs the application.
func (a *App) Run() {
	// Initialize termbox
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	termbox.SetOutputMode(termbox.Output256)
	defer termbox.Close()

	// Read configuration
	err = a.configuration.ReadConfiguration()
	if err != nil {
		log.Print("Unable to read configuration. Exiting.")
		return
	}

	// Read journal
	err = a.journal.ReadJournal()
	if err != nil {
		log.Print("Unable to read journal. Exiting.")
		return
	}

	// Initialize canvas
	width, height := io.GetWindowSize()

	// TODO: Use flex-box logic to size canvases
	bottomBarHeight := 7 // Height including borders
	mainViewport := screen.NewViewport(0, 0, width, height-bottomBarHeight)
	bottomViewport := screen.NewViewport(0, height-bottomBarHeight, width, bottomBarHeight)

	// Initialize screens
	a.configScreen = screens.NewConfigScreen(a.configuration, mainViewport)
	a.aboutScreen = screens.NewAboutScreen(a.configuration, mainViewport)
	a.journalScreen = screens.NewJournalScreen(a.configuration, mainViewport, a.journal)
	a.bottomBar = screens.NewBottomBarComponent(a.configuration, bottomViewport)

	// Register keypress handlers
	a.registerKeypressHandlers()

	// Render screen (initially)
	a.Render()

	// Start main app loop
	for a.isRunning {
		a.eventListener.WaitForEvent()
		a.Render()
	}
}

// Render ...
// Renders the current screen.
func (a *App) Render() {
	// Render current screen
	switch a.currentScreen {
	case ConfigScreen:
		a.configScreen.Render()
	case AboutScreen:
		a.aboutScreen.Render()
	case JournalScreen:
		a.journalScreen.Render()
	}

	// Render bottom bar
	a.bottomBar.Render()

	io.Flush()
}

// registerKeypressHandlers ...
// Registers the key press handlers.
func (a *App) registerKeypressHandlers() {
	// TODO: Screens should really register their own list of keys vs. having a single global list
	a.eventListener.RegisterKeypressHandler('?', a.showAboutScreen)
	a.eventListener.RegisterKeypressHandler('c', a.showConfigScreen)
	a.eventListener.RegisterKeypressHandler('j', a.showJournalScreen)
	a.eventListener.RegisterKeypressHandler('q', a.onQuit)
}

func (a *App) showConfigScreen() {
	a.currentScreen = ConfigScreen
}

func (a *App) showAboutScreen() {
	a.currentScreen = AboutScreen
}

func (a *App) showJournalScreen() {
	a.currentScreen = JournalScreen
}

// onQuit ...
// Called when the application should quit.
func (a *App) onQuit() {
	a.isRunning = false
}
