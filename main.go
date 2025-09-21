package main

import (
	"ashitomW/mazeGen/maze"
	"fmt"

	"github.com/alexflint/go-arg"
)

var args struct{
	Solve string `arg:"required"` 
}



func main() {

	arg.MustParse(&args)
	


	m := maze.NewMaze(15, 15)
	m.Generate(0, 0)

	var path []maze.Point

	switch args.Solve {
case "dfs":
		path = m.SolveDfs()	
	case "astar":
		path = m.SolveAStar()
	case "bfs":
		path = m.Solve()
	case "dijkstra":
		path = m.SolveDjkstra()
	}

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

