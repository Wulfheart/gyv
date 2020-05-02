package main

import (
	"github.com/Wulfheart/goyave-cli/cmd"
	"os"
)


func main() {
	// Dev
	err := os.Chdir("C:\\Users\\Alex\\Documents\\tmp")
	if err != nil {
		panic(err)
	}
	// Non-dev
	cmd.Run()
}
