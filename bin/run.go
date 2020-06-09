package main

import (
	"os"
	"quickstart"
)

func main() {
	proc := quickstart.NewHelloWorldProcess()
	proc.Run(os.Args)
}
