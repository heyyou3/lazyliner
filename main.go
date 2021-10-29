package main

import (
	"lazyliner/presentation"
	"os"
)

func main() {
	os.Exit(run())
}

func run() int {
	if err := presentation.ChooseCommandController(); err != nil {
		return 1
	}
	return 0
}
