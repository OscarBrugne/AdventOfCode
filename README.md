# Advent Of Code

Puzzles from [Advent of Code](https://adventofcode.com/) solved using Go.

My goal is to resolve problems on release day. Then, I refactor my code in order to have clean code, even if it means losing slightly in performance.

## Commands

- `cd <year>/<day>/` to navigate to a specific day's directory,
- `go run ./day<i>.go` to run the solution for a particular day,
- `go test` to run unit tests.
- `go test -bench .` to run benchmarks.

## 2023

Execution time with :
- goarch: `amd64`
- cpu: `AMD Ryzen 5 3500U with Radeon Vega Mobile Gfx`

| Day | Go | Part1     | Part2     | Comment                                                          |
|-----|----|-----------|-----------|------------------------------------------------------------------|
| 1   | ** | `0.039ms` | `1.51ms`  |                                                                  |
| 2   | ** | `0.44ms`  | `0.47ms`  |                                                                  |
| 3   | ** | `0.42ms`  | `0.55ms`  | Finding the numbers in the grid, then the symbols that are adjacent to them. |
| 4   | ** | `0.53ms`  | `0.52ms`  |                                                                  |
| 5   | ** | `0.12ms`  | `1.24ms`  | Using intervals, implementing the `splitOn` method to split the interval before shifting it. |
| 6   | ** | `0.002ms` | `47.2ms`  | Naive resolution, without calculating roots or binary search.    |
| 7   | ** | `7.35ms`  | `7.59ms`  | Using '*' instead of 'J' to represent jokers.                    |
| 8   | ** | `0.98ms`  | `3.96ms`  | Using LCM (Least Common Multiple) (doesn't work for general inputs, works here because the cycle length on each path is the same). |
| 9   | ** | `1.20ms`  | `1.21ms`  | Resolution without recursion.                                    |
| 10  | ** | `6.11ms`  | `30.6ms`  | Representation of a pipe by a set of directions. Using ray tracing for part 2. With visualization for part 2. |
| 11  | ** | `3.89ms`  | `3.80ms`  |                                                                  |
| 12  | ** | `11.8ms`  | `117ms`   | Using dynamic programming.                                       |
| 13  | ** | `0.30ms`  | `0.31ms`  | Using binary number to represent each line (in order to reduce execution time when comparing 2 lines). |
| 14  | ** | `2.30ms`  | `914ms`   | Using memoization to avoid redundant cycles.                     |
| 15  | ** | `0.14ms`  | `1.17ms`  |                                                                  |
| 16  | ** | `5.55ms`  | `1217ms`  | Part 2 in brute force. With visualization for part 1.            |
| 17  | ** | `519ms`   | `1844ms`  | Using A* algorithm (Implementation of IntPriorityQueue in utils). |
| 18  | ** | `0.11ms`  | `0.13ms`  | Using Shoelace formula and Pick's theorem.                       |
| 19  | ** | `1.07ms`  | `2.06ms`  | Parsing with `fmt.SscanfUsing`. Using intervals for part 2.      |
| 20  | *  | `0.00ms`  | `0.00ms`  | `TODO` (Part 2 not yet started)                                  |
| 21  | *  | `0.00ms`  | `0.00ms`  | `TODO` (Reflection done, implementation to be done)              |
| 22  |    | `0.00ms`  | `0.00ms`  | `TODO` (Not started yet)                                         |
| 23  | ** | `17.8ms`  | `1310ms`  | Transforming the grid into a graph (directed for part 1 and undirected for part 2) using Breadth First Search, and then getting the length of the longest simple path using Depth First Search. |
| 24  | *  | `0.00ms`  | `0.00ms`  | `TODO` (Part 2 : system of 6 equations with 6 unknowns found, I have to put it in matrix form and look for a solver in Go) |
| 25  | *  | `21.5ms`  |           | Separating the graph into 2 groups by calculating the shortest path between 2 vertices (with BFS), removing edges from this path, and repeating this 3 times (The minimum cut cardinality is 3 according to the puzzle). If there is no 4th path, the 2 vertices are in a different group and the new graph is cut, otherwise start again with a new vertex. Then finding the vertices belonging to the first group using BFS. |
