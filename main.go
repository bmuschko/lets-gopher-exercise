package main

import (
	"github.com/bmuschko/lets-gopher-exercise/cmd"
)

var (
	version = "undefined"
)

func main() {
	cmd.SetVersion(version)
	cmd.Execute()
}
