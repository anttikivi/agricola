package main

import (
	"os"

	"github.com/anttikivi/agricola/cmd"
)

func main() {
	os.Exit(cmd.Execute())
}
