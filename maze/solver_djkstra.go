package maze

import "math"


func (m *Maze) SolveDjkstra() []Point{
	start:= m.Grid[0][0]
	end:= m.Grid[m.Height-1][m.Width-1]


	dist := make(map[*Cell]int)
	prev := make(map[*Cell]*Cell)

	unvisited:= []*Cell{}

	for y:=0;y<m.Height;y++{
		for x:=0;x<m.Width;x++{
			cell:=m.Grid[y][x]
			dist[cell] = math.MaxInt
			unvisited = append(unvisited, cell)
		}
	}

	dist[start]=0
	for len(unvisited)>0{
		minIndex :=0
		for i, cell := range unvisited{
			if dist[cell] < dist[unvisited[minIndex]]{
				minIndex = i
			}
		}

		current := unvisited[minIndex]
		unvisited = append(unvisited[:minIndex],unvisited[minIndex+1:]...)

		if current == end {
			break
		}

		for _, neighbor := range m.getAccessibleNeighbors(current){
			alt := dist[current]+1
			if alt<dist[neighbor]{
				dist[neighbor]= alt
				prev[neighbor]=current
			}
		}

	}



	path := []Point{}
	for c:=end;c!=nil;c=prev[c]	{
		path = append([]Point{{c.X,c.Y}},path...)
	}
	return path
}
