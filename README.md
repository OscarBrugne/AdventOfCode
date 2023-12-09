# Advent Of Code

Puzzles from [Advent of Code](https://adventofcode.com/) solved using Go.

My goal is to resolve issues within the day of release. Then, I refactor my code in order to have clean code, even if it means losing slightly in performance.

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
| 1   | ** | `0.054ms` | `2.06ms`  |                                                                  |
| 2   | ** | `0.68ms`  | `0.75ms`  |                                                                  |
| 3   | ** | `0.42ms`  | `0.55ms`  | Finds the numbers in the grid, then the symbols that are adjacent to them. |
| 4   | ** | `0.53ms`  | `0.52ms`  |                                                                  |
| 5   | ** | `0.12ms`  | `1.24ms`  | Using intervals, implementing the `splitOn` method to split the interval before shifting it. |
| 6   | ** | `0.002ms` | `47.2ms`  | Naive resolution, without calculating roots or binary search.    |
| 7   | ** | `7.35ms`  | `7.59ms`  | Use '*' instead of 'J' to represent jokers.                      |
| 8   | ** | `0.98ms`  | `3.96ms`  | Using LCM (Least Common Multiple) (doesn't work for general inputs, works here because the cycle length on each path is the same). |
| 9   | ** | `1.20ms`  | `1.21ms`  |                                                                  |
