package consts

import (
	"os"
	"strings"
)

func Payment() []string {
	payments := os.Getenv("PAYMENTS")
	return strings.Split(payments, ",")
}