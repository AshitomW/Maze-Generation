package maze

import(
	"math/rand"
	"time"
)


func init(){
	rand.Seed(time.Now().UnixNano())
}

func (m *Maze) Generate(x,y int){
	stack:= []*Cell{}
	start:= m.Grid[y][x]
	start.visited = true
	stack = append(stack, start)


	for len(stack) > 0 {
		current := stack[len(stack)-1]
		neighbors := m.getUnvisitedNeighbors(current)
		if len(neighbors) > 0 {
			next:= neighbors[rand.Intn(len(neighbors))]
			m.removeWall(current,next)
			next.visited = true
			stack = append(stack, next)
		}else{
			//backtrack
			stack = stack[:len(stack)-1]
		}
	}

}



func (m *Maze) getUnvisitedNeighbors(c *Cell) []*Cell{
	neighbors := []*Cell{}
	// direction to move
	dirs := [][2]int{
		{0,-1}, // Top
		{1,0}, // Right
		{0,1}, // bottom
		{-1,0}, //left
	}

	for _,d := range dirs{
		// calculating neighbours coordinate
		nx, ny := c.x+d[0], c.y+d[1]
		if nx >= 0 && ny >=0 && nx <m.width && ny <m.height{
			neighbor := m.Grid	[ny][nx]
			if !neighbor.visited{
				neighbors = append(neighbors, neighbor)
			}
		}
	}

	return neighbors
}



// walls Index : 0-> Top, 1-> Right , 2-> Bottom ,  3-> Left
func (m *Maze) removeWall(a,b *Cell){
	dx := b.x - a.x
	dy := b.y - a.y 


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
