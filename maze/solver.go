package maze


type Point struct {x,y int}

type Node struct{
		Cell *Cell
		Parent *Node
	}

func (m *Maze) Solve() []Point{
	start := m.Grid[0][0]
	end := m.Grid[m.height-1][m.width-1]


	queue:= []*Node{{Cell: start}}
	visited:= make(map[*Cell]bool)

	visited[start] = true

	
	var endNode *Node

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]


		if current.Cell == end {
			endNode = current
			break
		}

		for _, neighbor := range m.getAccessibleNeighbors(current.Cell){
			if !visited[neighbor]{
				visited[neighbor] = true
				queue = append(queue, &Node{Cell: neighbor, Parent: current})

			}
		}

	}

	path := []Point{}
	for n:= endNode; n!=nil;n=n.Parent{
		path = append([]Point{{x: n.Cell.x, y:n.Cell.y}},path...)
	}

	return path

}



func (m *Maze) getAccessibleNeighbors(c *Cell) []*Cell{
	neighbors := []*Cell{}
	dirs:= [] struct{
		dx,dy int 
		wall int
		opp int
	}{
		{0, -1,0,2}, // Top
		{1,0,1,3}, // Right
		{0,1,2,0}, // Bottom
		{-1,0,3,1}, // Left
	}

	for _, d:= range dirs{
		nx, ny := c.x+d.dx, c.y+d.dy
		if nx>=0 && ny>=0 && nx<m.width && ny<m.height{
			neighbor:=m.Grid[ny][nx]
			if !c.Walls[d.wall] && !neighbor.Walls[d.opp]{
				neighbors = append(neighbors, neighbor)
			}
		}
	}
	return neighbors
}
