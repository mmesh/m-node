// Copyright (C) 2019 x6a
//
// pkg is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// pkg is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with pkg. If not, see <http://www.gnu.org/licenses/>.

package colors

import (
	"github.com/mgutz/ansi"
)

var Black = ansi.ColorFunc("black+bh")
var InvertedBlack = ansi.ColorFunc("white+bh:black")

var White = ansi.ColorFunc("white+bh")
var DarkWhite = ansi.ColorFunc("white+b")
var InvertedWhite = ansi.ColorFunc("black+b:white+h")

var Blue = ansi.ColorFunc("blue+bh")
var DarkBlue = ansi.ColorFunc("blue+b")
var InvertedBlue = ansi.ColorFunc("white+bh:blue")

var Cyan = ansi.ColorFunc("cyan+bh")
var DarkCyan = ansi.ColorFunc("cyan+b")
var InvertedCyan = ansi.ColorFunc("white+bh:cyan")

var Red = ansi.ColorFunc("red+bh")
var DarkRed = ansi.ColorFunc("red+b")
var InvertedRed = ansi.ColorFunc("white+bh:red")

var Green = ansi.ColorFunc("green+bh")
var DarkGreen = ansi.ColorFunc("green+b")
var InvertedGreen = ansi.ColorFunc("white+bh:green")

var Magenta = ansi.ColorFunc("magenta+bh")
var DarkMagenta = ansi.ColorFunc("magenta+b")
var InvertedMagenta = ansi.ColorFunc("white+bh:magenta")

var Yellow = ansi.ColorFunc("yellow+bh")
var DarkYellow = ansi.ColorFunc("yellow+b")
var InvertedYellow = ansi.ColorFunc("white+bh:yellow")
