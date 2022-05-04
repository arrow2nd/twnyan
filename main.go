package main

import (
	"github.com/arrow2nd/twnyan/cmd"
)

func main() {
	cmd := cmd.New()

	cmd.Init()
	cmd.Run()
}
