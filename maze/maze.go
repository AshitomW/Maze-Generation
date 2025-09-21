package maze


type Cell struct {
	x,y int
	Walls [4]bool 
	visited bool
}

type Maze struct{
	width,height int
	Grid [][]*Cell
}

func GenerateMaze(width,height int) *Maze{
	grid := make([][]*Cell, height)
	for y:=0;y<height;y++{
		grid[y] = make([]*Cell, width)
		for x:=0;x<width;x++{
			grid[y][x] = &Cell{
				x:x,y:y, Walls: [4]bool{true,true,true,true},
			}
		}
	}

	return &Maze{width:width,height:height,Grid: grid}
}
