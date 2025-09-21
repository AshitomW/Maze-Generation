package maze

import (
	"container/heap"
	"math"

)


type AStarNode struct{
	Cell *Cell
	G, H, F int // G Movement cost, H Heuristic Cost, F Total Cost
	Parent *AStarNode
	index int // for heap interface implementation
}


type PriorityQueue []*AStarNode

func (pq PriorityQueue) Len() int {return len(pq)}	
func (pq PriorityQueue) Less(i,j int) bool { return pq[i].F < pq[j].F}
func (pq PriorityQueue) Swap(i,j int) {
	pq[i],pq[j] = pq[j],pq[i];
	pq[i].index = i;
	pq[j].index = j;
}

func (pq *PriorityQueue) Push(x any){
	n:= x.(*AStarNode);
	n.index = len(*pq);
	*pq = append(*pq, n)
}

func (pq *PriorityQueue) Pop() any{
	old:= *pq;
	n:=len(old);
	node:= old[n-1];
	*pq = old[0:n-1];
	return node
} 

func (m *Maze) SolveAStar() []Point {

	
	start:= m.Grid[0][0]
	end:=m.Grid[m.Height-1][m.Width-1]

	openSet := &PriorityQueue{}
	heap.Init(openSet)

	startNode := &AStarNode{Cell: start, G:0,H:manhattanDistance(start,end)}
	startNode.F = startNode.G + startNode.H
	heap.Push(openSet,startNode)

	gScore:=make(map[*Cell]int)
	gScore[start]=0


	for openSet.Len() > 0 {
		current := heap.Pop(openSet).(*AStarNode)
		if current.Cell == end {
			return reconstructPath(current)
		}
		for _, neighbor := range m.getAccessibleNeighbors(current.Cell){
			tentativeG := current.G + 1
			if g,ok := gScore[neighbor]; !ok || tentativeG < g{
				gScore[neighbor] = tentativeG
				neighborNode := &AStarNode{
					Cell: neighbor,
					G: tentativeG,
					H: manhattanDistance(neighbor,end),
					Parent: current,
				}
				neighborNode.F = neighborNode.G + neighborNode.H 
				heap.Push(openSet,neighborNode)
			}
		}
	}


	return nil
}




func manhattanDistance(a,b *Cell) int {
	return int(math.Abs(float64(a.X-b.X))+math.Abs(float64(a.Y-b.Y)))
}

func reconstructPath(node *AStarNode) []Point{
	path:= []Point{}
	for n:=node;n!=nil;n=n.Parent{
		path = append([]Point{{n.Cell.X,n.Cell.Y}},path...)
	}

	return path
}
