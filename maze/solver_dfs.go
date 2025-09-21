package maze


func (m *Maze) SolveDfs() []Point{
	visited := make(map[*Cell] bool)
	path := []Point{}
	if m.dfs(m.Grid[0][0],m.Grid[m.Height-1][m.Width-1],visited,&path){
		return path
	}
	return nil
}



func (m *Maze) dfs(current,end *Cell, visited map[*Cell]bool, path *[]Point) bool{
	visited[current] = true
	*path = append(*path, Point{current.X,current.Y})

	if current == end{
		return true // exit found
	}


	for _,neighbor := range m.getAccessibleNeighbors(current){
		if !visited[neighbor]{
			if m.dfs(neighbor,end,visited,path){
				return true // path is found through a neighbor
			}
		}
	}

	// backtracking on deadend

	*path = (*path)[:len(*path)-1]
	return  false
	
}
