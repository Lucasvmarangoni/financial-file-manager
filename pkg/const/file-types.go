package consts

import (
	"os"
	"strings"
)

func FileTypes() []string {
	types := os.Getenv("FILE_TYPES")
	return strings.Split(types, ",")
}