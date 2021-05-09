package main

import (
	"fmt"
	"os"

	"github.com/miketmoore/mazegen"
)

func main() {
	rows := 10
	cols := 10
	verbose := false
	grid, err := mazegen.BuildMaze(rows, cols, verbose)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	fmt.Println(grid)
}
