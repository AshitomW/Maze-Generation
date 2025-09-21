package main

import (
	"fmt"
	"ashitomW/mazeGen/maze"
)

func main() {
	m := maze.NewMaze(10, 10)
	m.Generate(0, 0)

	path := m.Solve()

	for y := 0; y < m.Height; y++ {
			for x := 0; x < m.Width; x++ {
			if m.Grid[y][x].Walls[0] {
				fmt.Print("+---")
			} else {
				fmt.Print("+   ")
			}
		}
		fmt.Println("+")

		for x := 0; x < m.Width; x++ {
			if m.Grid[y][x].Walls[3] {
				fmt.Print("|")
			} else {
				fmt.Print(" ")
			}
			inPath := false
			for _, p := range path {
				if p.X == x && p.Y == y {
					inPath = true
					break
				}
			}

 if x == 0 && y == 0 {
    fmt.Print(" S ") // Start
 } else if x == m.Width-1 && y == m.Height-1 {
    fmt.Print(" E ") // End
 } else if inPath {
    fmt.Print(" * ") 
 } else {
    fmt.Print("   ")}

		}

		fmt.Println("|")
	}

	for x := 0; x < m.Width; x++ {
		fmt.Print("+---")
	}
	fmt.Println("+")
}

