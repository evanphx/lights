package main

import (
	"github.com/evanphx/lights"
	"github.com/jessevdk/go-flags"
)

func main() {
	parser := flags.NewParser(&lights.Globals, flags.Default)

	parser.AddCommand("off", "turn a light off", "", &lights.OffCommand{})
	parser.AddCommand("on", "turn a light on", "", &lights.OnCommand{})
	parser.AddCommand("list", "list all the lights", "", &lights.ListCommand{})
	parser.AddCommand("register", "register with the bridge", "", &lights.RegisterCommand{})

	parser.Parse()
}
