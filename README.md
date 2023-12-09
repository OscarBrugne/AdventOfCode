# Advent Of Code

Puzzles from [Advent of Code](https://adventofcode.com/) solved using Go.

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
| 3   | ** | `0.26ms`  | `0.42ms`  |                                                                  |
| 4   | ** | `0.79ms`  | `0.71ms`  |                                                                  |
| 5   | ** | `0.13ms`  | `0.27ms`  | Calculation on intervals, function "splitOverlappingIntervals" to refactor |
| 6   | ** | `0.002ms` | `66.2ms`  | Naive resolution, without calculating roots.                     |
| 7   | ** | `10.2ms`  | `12.8ms`  |                                                                  |
| 8   | ** | `0.98ms`  | `5.06ms`  | Using LCM (least common multiple) (doesn't work for general inputs, works here because the cycle length on each path is the same.) |
| 9   | ** | `1.20ms`  | `1.21ms`  |                                                                  |
