package maze

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type wall struct {
	x1, y1 int
	x2, y2 int
}

// Using randomized Prim's algorithm with loop bias
func (m *Maze) Generate(x, y int) {

	for row := range m.Grid {
		for _, c := range m.Grid[row] {
			c.visited = false
			c.Walls = [4]bool{true, true, true, true}
		}
	}

	start := m.Grid[y][x]
	start.visited = true

	frontier := []wall{}
	m.addFrontierWalls(start, &frontier)

	for len(frontier) > 0 {
		idx := rand.Intn(len(frontier))
		w := frontier[idx]

		frontier[idx] = frontier[len(frontier)-1]
		frontier = frontier[:len(frontier)-1]

		a := m.Grid[w.y1][w.x1]
		b := m.Grid[w.y2][w.x2]

		//  1: one side visited, the other unvisited → normal carve
		if a.visited != b.visited {
			m.removeWall(a, b)
			if !a.visited {
				a.visited = true
				m.addFrontierWalls(a, &frontier)
			}
			if !b.visited {
				b.visited = true
				m.addFrontierWalls(b, &frontier)
			}
		} else {
			// 2: both already visited → loop bias , 5% chance of removing wall when both side are already visited
			if rand.Float64() < 0.05 {
				m.removeWall(a, b)
			}
		}
	}
}


func (m *Maze) addFrontierWalls(c *Cell, frontier *[]wall) {
	dirs := [][2]int{
		{0, -1}, // top
		{1, 0},  // right
		{0, 1},  // bottom
		{-1, 0}, // left
	}
	for _, d := range dirs {
		nx, ny := c.X+d[0], c.Y+d[1]
		if nx >= 0 && ny >= 0 && nx < m.Width && ny < m.Height {
			*frontier = append(*frontier, wall{c.X, c.Y, nx, ny})
		}
	}
}


func (m *Maze) removeWall(a,b *Cell){
	dx := b.X - a.X
	dy := b.Y - a.Y 


	// remove wall from a and b of index specified 
	switch {
		case dx == 1: // b is in the right of a
			a.Walls[1] = false 
			b.Walls[3] = false 
		case  dx == -1: 
			a.Walls[3] = false 
			b.Walls[1] = false 
	case dy == 1: 
		a.Walls[2] = false
		b.Walls[0] = false
	case dy == -1:
		a.Walls[0] = false
		b.Walls[2] = false
	}
}

