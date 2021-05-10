package main

import (
	"fmt"
	"os"

	"github.com/miketmoore/mazegen"
)

func main() {
	rows := 2
	cols := 2
	random := mazegen.NewRandom()
	grid, err := mazegen.BuildMaze(rows, cols, random)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	for rowIndex, row := range grid.Cells {
		for columnIndex, cell := range row {
			fmt.Println(rowIndex, columnIndex, cell)
		}
	}
}
