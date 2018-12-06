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
	"github.com/stuartthompson/fitnesse-term/io/screen"
)

// ConfigScreen ...
type ConfigScreen struct {
	screen        *screen.Screen
	configuration *configuration.AppConfig
}

// NewConfigScreen ...
// Instantiates a new config screen.
func NewConfigScreen(config *configuration.AppConfig, viewport *screen.Viewport) *ConfigScreen {
	screenStyle := &screen.Style{ShowBorder: true, BorderColor: 100}
	screen := screen.NewScreen(viewport, screenStyle)
	return &ConfigScreen{screen: screen, configuration: config}
}

// Render ...
// Renders the config screen.
func (s *ConfigScreen) Render() {
	s.screen.Clear()

	s.screen.RenderText("Config", 1, 1, 255, 0)
	s.screen.RenderText("Default language: "+s.configuration.DefaultLanguage, 1, 3, 255, 0)
}
