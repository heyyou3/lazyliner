package main

import (
	"fmt"
	"lazyliner/infra/command"
	"os"
)

func main() {
	config, err := command.NewTomlRepository("./config.toml")
	if err != nil {
		os.Exit(1)
	}
	cmd := config.Fetch()
	fmt.Println(cmd)
	fmt.Println(run())
}

func run() int {
	return 0
}
