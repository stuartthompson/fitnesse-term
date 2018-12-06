// Copyright 2018 Stuart Thompson.

// This file is part of Fitnesse-Term

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

// borderColor ...
// The border color for this component.
const borderColor = 54 // Purple (#5f0087)

// BottomBarComponent ...
type BottomBarComponent struct {
	screen *screen.Screen
	config *configuration.AppConfig
}

// NewBottomBarComponent ...
// Instantiates a new bottom bar component.
func NewBottomBarComponent(config *configuration.AppConfig, viewport *screen.Viewport) *BottomBarComponent {
	screenStyle := &screen.Style{ShowBorder: true, BorderColor: borderColor}
	screen := screen.NewScreen(viewport, screenStyle)
	return &BottomBarComponent{screen: screen, config: config}
}

// Render ...
// Renders the bottom bar component.
func (c *BottomBarComponent) Render() {
	c.screen.Clear()

	c.screen.RenderText("Bottom Bar", 0, 0, 255, 0)
}
