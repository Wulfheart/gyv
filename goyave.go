package main

import (
	"github.com/Wulfheart/goyave-cli/cmd"
	"os"
)

func main() {
	// For local testing
	os.Chdir("C:\\Users\\Alex\\Documents\\tmp")
	cmd.Run()
}
