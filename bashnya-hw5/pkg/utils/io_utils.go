package io_utils

import (
	"fmt"
	"os"
)

func PrintError(err error) {
	fmt.Fprintln(os.Stderr, err)
}
