package colors

import (
	"github.com/mgutz/ansi"
)

var Black = ansi.ColorFunc("black+bh")
var InvertedBlack = ansi.ColorFunc("black+b:black+h")

var InvertedBlackWhite = ansi.ColorFunc("white+b:black+h")

var White = ansi.ColorFunc("white+bh")
var DarkWhite = ansi.ColorFunc("white+b")
var InvertedWhite = ansi.ColorFunc("black+b:white+h")
var InvertedDarkWhite = ansi.ColorFunc("black+b:white")

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
