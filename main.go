package main

import (
	"ashitomW/mazeGen/maze"
	"fmt"
)

func main() {
	m := maze.NewMaze(15, 15)
	m.Generate(0, 0)

	results := map[string][]maze.Point{
		"DFS":      m.SolveDfs(),
		"BFS":      m.Solve(),
		"Dijkstra": m.SolveDjkstra(),
		"A*":       m.SolveAStar(),
	}


	for name, path := range results {
		fmt.Printf("\n=== %s Solution ===\n\n", name)
		printMazeWithPath(m, path)
	}
}

func printMazeWithPath(m *maze.Maze, path []maze.Point) {
	for y := 0; y < m.Height; y++ {
		// Top walls
		for x := 0; x < m.Width; x++ {
			if m.Grid[y][x].Walls[0] {
				fmt.Print("+---")
			} else {
				fmt.Print("+   ")
			}
		}
		fmt.Println("+")

		// Cells and vertical walls
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
				fmt.Print(" X ")
			} else {
				fmt.Print("   ")
			}
		}
		fmt.Println("|")
	}

	// Bottom border
	for x := 0; x < m.Width; x++ {
		fmt.Print("+---")
	}
	fmt.Println("+")
}

