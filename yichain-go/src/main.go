package main

import (
	"fmt"
	"os"
	"yichain-go/app"
)

func main() {
	if err := app.Run(os.Args); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

