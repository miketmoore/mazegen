package main

import "github.com/miketmoore/mazegen"

func main() {
	rows := 10
	cols := 10
	verbose := false
	mazegen.BuildMaze(rows, cols, verbose)
}
