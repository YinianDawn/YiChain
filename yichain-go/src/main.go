package main

import (
	"os"
	"yichain-go/app"
)

func main() {
	if config, err := app.Run(os.Args); err != nil {
		if config != nil {
			config.Log.Error(err.Error())
		}
		os.Exit(1)
	}
}
