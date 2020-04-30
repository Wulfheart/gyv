package cmd

import "github.com/leaanthony/clir"

var cli = clir.NewCli("goyave", "The go web framework with a special flavour.","v0.1.0")



func init(){
	cli.AddCommand(createCmd)
}

//Run starts the cli
func Run(){
	if err := cli.Run(); err != nil{
		panic(err)
	}
}
