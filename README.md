# Maze Generator and Solver

A Go-based maze generation and pathfinding project that implements multiple maze solving algorithms including A\*, Dijkstra's, Depth-First Search (DFS), and Breadth-First Search (BFS).

## Features

- **Maze Generation**: Creates random mazes using randomized Prim's algorithm with loop bias for maze structures
- **Multiple Solving Algorithms**:
  - **A\*** - Heuristic-based pathfinding with Manhattan distance
  - **Dijkstra's** - Shortest path algorithm for weighted graphs
  - **BFS** - Breadth-first search for unweighted shortest path
  - **DFS** - Depth-first search with backtracking
- **Visual Output**: ASCII representation of the maze with solution paths marked for all algorithms

## Project Structure

```
mazeGen/
├── main.go              # Main entry point with visualization
├── go.mod              # Go module definition
├── README.md           # This file
└── maze/
    ├── maze.go         # Core maze structure and creation
    ├── generator.go    # Randomized Prim's maze generation
    ├── solver_bfs.go   # Breadth-first search implementation
    ├── solver_dfs.go   # Depth-first search implementation
    ├── solver_djkstra.go # Dijkstra's algorithm implementation
    └── solver_astar.go # A* algorithm implementation
```

## Installation

1. Clone the repository:

```bash
git clone <repository-url>
cd mazeGen
```

2. Run the project:

```bash
go run main.go
```

Or build and run:

```bash
go build -o mazeGen main.go
./mazeGen
```

## Usage

Simply run the program to generate a 15x15 maze and see all four solving algorithms in action:

```bash
go run main.go
```

The program will:

1. Generate a random 15x15 maze using randomized Prim's algorithm
2. Solve the maze using all four algorithms (DFS, BFS, Dijkstra, A\*)
3. Display each solution with the path marked

## Algorithm Details

### Maze Generation

- **Randomized Prim's Algorithm** with 5% loop bias
- Creates more interesting mazes with occasional loops
- Starts from position (0,0) and grows the maze randomly

### Solving Algorithms

- **DFS (Depth-First Search)**: Uses recursion and backtracking to find a path
- **BFS (Breadth-First Search)**: Guarantees shortest path in unweighted graphs
- **Dijkstra's Algorithm**: Finds shortest path using distance-based exploration
- **A\* Algorithm**: Uses Manhattan distance heuristic for efficient pathfinding

## Configuration

The maze size is set to 15x15 in `main.go`. To change dimensions, modify:

```go
m := maze.NewMaze(15, 15)  // width, height
```
