package logger

import (
	"strings"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"

	"os"
)

func init() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, NoColor: false, TimeFormat: time.RFC3339,
		FormatMessage: func(i interface{}) string {
			value, ok := i.(string)
			if !ok {
				return "unknown"
			}
			return Format(c.magenta, strings.ToUpper(value))
		},
		FormatLevel: func(i interface{}) string {
			level, ok := i.(string)
			if !ok {
				return "unknown"
			}
			switch level {
			case "info":
				return Format(c.green, strings.ToUpper(level)+" ⇝")
			case "error":
				return Format(c.yellow, strings.ToUpper(level)+" ⇝")
			case "fatal":
				return Format(c.red, strings.ToUpper(level)+" ⇝")
			default:
				return level
			}
		},
		FormatErrFieldValue: func(i interface{}) string {
			value, ok := i.(string)
			if !ok {
				return "unknown"
			}
			formattedOperation := Format(c.blue, "Operation")
			formattedError := Format(c.red, "Error")
			Str := strings.ReplaceAll(value, "Operation", formattedOperation)
			Str = strings.ReplaceAll(Str, "Error", formattedError)
			return Str
		},
		FormatErrFieldName: func(i interface{}) string {
			_, ok := i.(string)
			if !ok {
				return "unknown"
			}
			return " "
		},
	})
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
}
