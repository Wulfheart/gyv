package cmd

import (
	"fmt"
	"github.com/leaanthony/clir"
)

var createCmd = clir.NewCommand("create", "Creates an empty project")


func init(){
	var stringer string
	//createCmd.StringFlag("sdfg", "Test", &stringer)
	createCmd.Action(func() error{
		fmt.Println(stringer)
		return nil
	})
}