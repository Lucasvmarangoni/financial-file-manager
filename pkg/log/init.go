package logger

import (
	"fmt"
	"strings"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"

	"os"
)

func init() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout, NoColor: false, TimeFormat: time.RFC3339,
		FormatMessage: func(i interface{}) string {
			value := failOnError(i)
			return Format(colorMagenta, strings.ToUpper(value))
		},
		FormatLevel: func(i interface{}) string {
			level := failOnError(i)
			switch level {
			case "info":
				return Format(colorGreen, strings.ToUpper(level)+" ⇝")
			case "error":
				return Format(colorRed, strings.ToUpper(level)+" ⇝")
			case "warn":
				return Format(colorYellow, strings.ToUpper(level)+" ⇝")
			case "debug":
				return Format(colorCyan, strings.ToUpper(level)+" ⇝")
			case "fatal":
				return Format(colorRed, strings.ToUpper(level)+" ⇝")
			default:
				return level
			}
		},
		FormatErrFieldValue: func(i interface{}) string {
			value := failOnError(i)
			formattedOperation := Format(colorBlue, "Operation")
			formattedError := Format(colorRed, "Error")
			Str := strings.ReplaceAll(value, "Operation", formattedOperation)
			Str = strings.ReplaceAll(Str, "Error", formattedError)
			return Str
		},
		FormatErrFieldName: func(i interface{}) string {
			value := failOnError(i)
			return value
		},
	})
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	zerolog.CallerMarshalFunc = func(pc uintptr, file string, line int) string {
		return fmt.Sprintf("[\n%s:%d\n]", file, line)
	}
}

func failOnError(i interface{}) string {
	value, ok := i.(string)
	if !ok {
		return "unknown"
	}
	return value
}
