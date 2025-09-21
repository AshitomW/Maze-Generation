package maze


type Cell struct {
	X,Y int
	Walls [4]bool 
	Visited bool
}

type Maze struct{
	Width,Height int
	Grid [][]*Cell
}

func NewMaze(width,height int) *Maze{
	grid := make([][]*Cell, height)
	for y:=0;y<height;y++{
		grid[y] = make([]*Cell, width)
		for x:=0;x<width;x++{
			grid[y][x] = &Cell{
				X:x,Y:y, Walls: [4]bool{true,true,true,true},
			}
		}
	}

	return &Maze{Width:width,Height:height,Grid: grid}
}
