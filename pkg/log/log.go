package logger

import (
	"fmt"

	"github.com/fatih/color"
)

const (
	colorBlack = iota + 30
	colorRed
	colorGreen
	colorYellow
	colorBlue
	colorMagenta
	colorCyan
	colorWhite
)

func Config() {}

func Format(col color.Attribute, value string) string {
	return fmt.Sprintf("\x1b[%dm%v\x1b[0m", col, value)
}
