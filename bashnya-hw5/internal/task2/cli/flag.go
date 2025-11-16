package cli

import (
	"flag"
	"fmt"
)

func ParseFlags() (int, error) {
	workers := flag.Int("workers", 1, "number of workers")
	flag.Parse()

	if *workers <= 0 {
		return 0, fmt.Errorf("invalid num of workers")
	}

	return *workers, nil
}
