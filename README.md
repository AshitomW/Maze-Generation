# Maze Generator and Solver

A Go-based maze generation and pathfinding project that implements multiple maze solving algorithms including A\*, Dijkstra's, Depth-First Search (DFS), and Breadth-First Search (BFS).

## Features

- **Maze Generation**: Creates random mazes using a depth-first search algorithm with backtracking
- **Multiple Solving Algorithms**:
  - **A\*** - Heuristic-based pathfinding with Manhattan distance
  - **Dijkstra's** - Shortest path algorithm for weighted graphs
  - **BFS** - Breadth-first search for unweighted shortest path
  - **DFS** - Depth-first search with backtracking
- **Visual Output**: ASCII representation of the maze with the solution path marked

## Installation

1. Clone the repository:

```bash
git clone <repository-url>
cd mazeGen
```

2. Install dependencies:

```bash
go mod tidy
```

3. Build the project:

```bash
go build -o mazeGen main.go
```

## Usage

Run the maze generator and solver with the following command:

```bash
./mazeGen --solve <algorithm>
```

### Available Algorithms

- `astar` - A\* pathfinding algorithm
- `dijkstra` - Dijkstra's shortest path algorithm
- `bfs` - Breadth-first search
- `dfs` - Depth-first search

### Examples

```bash
# Solve using A* algorithm
./mazeGen --solve astar

# Solve using Dijkstra's algorithm
./mazeGen --solve dijkstra

# Solve using BFS
./mazeGen --solve bfs

# Solve using DFS
./mazeGen --solve dfs
```
Since the maze is generated randomly each time, the differences between algorithms is not  visually obvious. However, each algorithm follows its own unique search strategy, which affects path length and exploration patterns.

## Output

The program generates a 15x15 maze and displays it in ASCII format:

- `S` marks the start position (top-left)
- `E` marks the end position (bottom-right)
- `*` marks the solution path
- `+`, `-`, `|` represent maze walls

Example output:

```
+---+---+---+---+---+
| S     |           |
+   +---+   +---+   +
|   |       |   | * |
+   +   +---+   +   +
|       |       | * |
+---+---+   +---+   +
|           |     E |
+---+---+---+---+---+
```

## Dependencies

- [github.com/alexflint/go-arg](https://github.com/alexflint/go-arg) - Command line argument parsing

## Configuration

The maze size is currently hardcoded to 15x15 in `main.go`. You can modify the dimensions by changing:

```go
m := maze.NewMaze(15, 15)  // width, height
```
