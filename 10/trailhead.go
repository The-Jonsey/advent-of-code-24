package main

import (
	"advent24/common"
	"strconv"
)

func main() {
	grid := common.InputTo2dArray(common.GetInput("10"))
	trailheadScores := 0
	for y := range grid {
		for x, col := range grid[y] {
			if col == "0" {
				trailheads := make(map[string]int)
				trailheads = GetTrailheads(0, x, y, grid, trailheads)
				//trailheadScores += len(trailheads) -- for part 1
				for _, i := range trailheads {
					trailheadScores += i
				}
			}
		}
	}
	println(trailheadScores)
}

func GetTrailheads(currentNumber int, x int, y int, grid [][]string, trailheads map[string]int) map[string]int {
	if currentNumber == 9 {
		seenCount, ok := trailheads[strconv.Itoa(x)+","+strconv.Itoa(y)]
		if !ok {
			seenCount = 0
		}
		trailheads[strconv.Itoa(x)+","+strconv.Itoa(y)] = seenCount + 1
		return trailheads
	}

	nextNumber := currentNumber + 1
	nextNumberString := strconv.Itoa(nextNumber)

	if x > 0 && grid[y][x-1] == nextNumberString {
		GetTrailheads(nextNumber, x-1, y, grid, trailheads)
	}

	if y > 0 && grid[y-1][x] == nextNumberString {
		GetTrailheads(nextNumber, x, y-1, grid, trailheads)
	}

	if x+1 < len(grid[0]) && grid[y][x+1] == nextNumberString {
		GetTrailheads(nextNumber, x+1, y, grid, trailheads)
	}

	if y+1 < len(grid) && grid[y+1][x] == nextNumberString {
		GetTrailheads(nextNumber, x, y+1, grid, trailheads)
	}

	return trailheads
}
