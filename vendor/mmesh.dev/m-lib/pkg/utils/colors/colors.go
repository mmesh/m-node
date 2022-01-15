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
var InvertedBlue = ansi.ColorFunc("white+bh:blue+h")
var InvertedBlueAlt = ansi.ColorFunc("black+b:blue+h")
var InvertedDarkBlue = ansi.ColorFunc("white+bh:blue")
var InvertedDarkBlueAlt = ansi.ColorFunc("black+b:blue")

var Cyan = ansi.ColorFunc("cyan+bh")
var DarkCyan = ansi.ColorFunc("cyan+b")
var InvertedCyan = ansi.ColorFunc("white+bh:cyan+h")
var InvertedCyanAlt = ansi.ColorFunc("black+b:cyan+h")
var InvertedDarkCyan = ansi.ColorFunc("white+bh:cyan")
var InvertedDarkCyanAlt = ansi.ColorFunc("black+b:cyan")

var Red = ansi.ColorFunc("red+bh")
var DarkRed = ansi.ColorFunc("red+b")
var InvertedRed = ansi.ColorFunc("white+bh:red+h")
var InvertedRedAlt = ansi.ColorFunc("black+b:red+h")
var InvertedDarkRed = ansi.ColorFunc("white+bh:red")
var InvertedDarkRedAlt = ansi.ColorFunc("black+b:red")

var Green = ansi.ColorFunc("green+bh")
var DarkGreen = ansi.ColorFunc("green+b")
var InvertedGreen = ansi.ColorFunc("white+bh:green+h")
var InvertedGreenAlt = ansi.ColorFunc("black+b:green+h")
var InvertedDarkGreen = ansi.ColorFunc("white+bh:green")
var InvertedDarkGreenAlt = ansi.ColorFunc("black+b:green")

var Magenta = ansi.ColorFunc("magenta+bh")
var DarkMagenta = ansi.ColorFunc("magenta+b")
var InvertedMagenta = ansi.ColorFunc("white+bh:magenta+h")
var InvertedMagentaAlt = ansi.ColorFunc("black+b:magenta+h")
var InvertedDarkMagenta = ansi.ColorFunc("white+bh:magenta")
var InvertedDarkMagentaAlt = ansi.ColorFunc("black+b:magenta")

var Yellow = ansi.ColorFunc("yellow+bh")
var DarkYellow = ansi.ColorFunc("yellow+b")
var InvertedYellow = ansi.ColorFunc("white+bh:yellow+h")
var InvertedYellowAlt = ansi.ColorFunc("black+b:yellow+h")
var InvertedDarkYellow = ansi.ColorFunc("white+bh:yellow")
var InvertedDarkYellowAlt = ansi.ColorFunc("black+b:yellow")
