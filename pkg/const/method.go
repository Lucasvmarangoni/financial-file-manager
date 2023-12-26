package consts

import (
	"os"
	"strings"
)

func Method() []string {
	methods := os.Getenv("METHODS")
	return strings.Split(methods, ",")
}
