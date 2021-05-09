package main

import (
	"fmt"
	"os"

	"github.com/miketmoore/mazegen"
)

func main() {
	rows := 10
	cols := 10
	random := mazegen.NewRandom()
	grid, err := mazegen.BuildMaze(rows, cols, random)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	fmt.Println(grid)
}
