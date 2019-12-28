package main

import (
	"fmt"
	"github.com/YinianDawn/yichain/yichain-go/src/app"
	"os"
)

func main() {
	if err := app.Run(os.Args); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

