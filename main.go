package main

import (
	"fmt"
	"github.com/amiranbari/challenge/config"
)

func main() {
	cfg := Config()
	fmt.Println(cfg)
}

func Config() config.Config {
	return config.C()
}
