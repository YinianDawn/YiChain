package app

import (
	"fmt"
	"yichain-go/config"
)

func Run(args []string) (*config.Config, error) {
	cfg := config.New()
	fmt.Println(args)
	cfg.Log.Trace("123 {}", "trace")
	cfg.Log.Info("123 {}", 1)
	cfg.Log.Warning("123 {}", "213231")
	cfg.Log.Error("123 {}", 1.1)
	return cfg, nil
}
