package utils

import (
	"fmt"
	"github.com/pkg/errors"
	"os"
)

func CheckIfError(err error) {
	if err == nil {
		return
	}

	fmt.Printf("\x1b[31;1m%s\x1b[0m\n", errors.WithMessage(err, "error"))
	os.Exit(1)
}
