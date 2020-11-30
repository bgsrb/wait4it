package main

import (
	"wait4it/cmd"
	"wait4it/config"
)

func main() {
	cmd.Run(config.Parse())
}
