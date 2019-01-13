package main

import (
	"github.com/bmuschko/lets-gopher/cmd"
)

var (
	version = "undefined"
)

func main() {
	cmd.SetVersion(version)
	cmd.Execute()
}