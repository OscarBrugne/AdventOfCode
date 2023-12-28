package main

import (
	"fmt"
	"sort"
	"time"

	"AdventOfCode/utils"
)

type Coord struct {
	x int
	y int
	z int
}

type Brick struct {
	mini    Coord
	maxi    Coord
	basedOn []*Brick
	support []*Brick
}

func parseInput(input []string) []*Brick {
	bricks := make([]*Brick, len(input))
	for i, line := range input {
		brick := Brick{
			basedOn: []*Brick{},
			support: []*Brick{},
		}
		fmt.Sscanf(line, "%d,%d,%d~%d,%d,%d", &brick.mini.x, &brick.mini.y, &brick.mini.z, &brick.maxi.x, &brick.maxi.y, &brick.maxi.z)
		bricks[i] = &brick
	}
	return bricks
}

func settle(bricks []*Brick) {
	sort.Slice(bricks, func(i, j int) bool {
		return bricks[i].maxi.z < bricks[j].maxi.z
	})

	for i, brick := range bricks {
		supportZ := 0
		basedBricks := []*Brick{}

		for j := i - 1; j > -1; j-- {
			isIntersectingX := max(brick.mini.x, bricks[j].mini.x) <= min(brick.maxi.x, bricks[j].maxi.x)
			isIntersectingY := max(brick.mini.y, bricks[j].mini.y) <= min(brick.maxi.y, bricks[j].maxi.y)
			isIntersecting := isIntersectingX && isIntersectingY
			if isIntersecting {
				if bricks[j].maxi.z == supportZ {
					basedBricks = append(basedBricks, bricks[j])
				} else if bricks[j].maxi.z > supportZ {
					supportZ = bricks[j].maxi.z
					basedBricks = []*Brick{bricks[j]}
				}
			}
		}

		brick.basedOn = basedBricks
		for _, basedBrick := range basedBricks {
			basedBrick.support = append(basedBrick.support, brick)
		}

		deltaZ := brick.maxi.z - brick.mini.z
		brick.mini.z = supportZ + 1
		brick.maxi.z = brick.mini.z + deltaZ
	}
}

func Part1(input []string) int {
	bricks := parseInput(input)
	settle(bricks)

	cnt := 0
	for _, brick := range bricks {
		isDisintegratable := true
		for _, supportedBrick := range brick.support {
			if len(supportedBrick.basedOn) < 2 {
				isDisintegratable = false
				break
			}
		}
		if isDisintegratable {
			cnt++
		}
	}
	return cnt
}

func Part2(input []string) int {
	bricks := parseInput(input)
	settle(bricks)

	cnt := 0
	for _, brick := range bricks {
		fallingBricks := map[*Brick]struct{}{}
		for _, supportedBrick := range brick.support {
			if len(supportedBrick.basedOn) == 1 {
				allSupportedBricks := []*Brick{supportedBrick}
				for len(allSupportedBricks) > 0 {
					supportedBrick0 := allSupportedBricks[0]
					allSupportedBricks = allSupportedBricks[1:]

					isFalling := true
					for _, basedBrick := range supportedBrick0.basedOn {
						if _, ok := fallingBricks[basedBrick]; basedBrick != brick && !ok {
							isFalling = false
							break
						}
					}

					if isFalling {
						fallingBricks[supportedBrick0] = struct{}{}
						allSupportedBricks = append(allSupportedBricks, supportedBrick0.support...)
					}
				}
			}
		}
		cnt += len(fallingBricks)
	}
	return cnt
}

func main() {
	fileName := "input.txt"
	input := utils.ReadFile(fileName)

	start1 := time.Now()
	fmt.Println("Answer 1 : ", Part1(input))
	fmt.Println(time.Since(start1))

	start2 := time.Now()
	fmt.Println("Answer 2 : ", Part2(input))
	fmt.Println(time.Since(start2))
}
