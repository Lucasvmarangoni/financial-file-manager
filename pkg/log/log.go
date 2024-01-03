package logger

import (
	"github.com/fatih/color"
)

type colored = color.Attribute

type colors struct {
	green   colored
	yellow  colored
	red     colored
	blue    colored
	cyan    colored
	magenta colored
}

var c = &colors{
	green:   color.FgGreen,
	yellow:  color.FgYellow,
	red:     color.FgRed,
	blue:    color.FgBlue,
	cyan:    color.FgCyan,
	magenta: color.FgMagenta,
}

func Config() *colors {
	return c
}

func Format(col color.Attribute, value string) string {
	return color.New(col).Sprint(value)
}
