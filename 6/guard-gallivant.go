package main

import (
	"advent24/common"
)

type GuardDirection interface {
	transformGuardPos(x int, y int) []int
	getName() string
}

type UpwardsGuardDirection struct {
}

func (UpwardsGuardDirection) transformGuardPos(x int, y int) []int {
	return []int{x, y - 1}
}

type DownwardsGuardDirection struct {
}

func (DownwardsGuardDirection) transformGuardPos(x int, y int) []int {
	return []int{x, y + 1}
}

type LeftGuardDirection struct {
}

func (LeftGuardDirection) transformGuardPos(x int, y int) []int {
	return []int{x - 1, y}
}

type RightGuardDirection struct {
}

func (RightGuardDirection) transformGuardPos(x int, y int) []int {
	return []int{x + 1, y}
}

var UPWARDS = UpwardsGuardDirection{}
var LEFT = LeftGuardDirection{}
var RIGHT = RightGuardDirection{}
var DOWN = DownwardsGuardDirection{}

func nextGuardDirection(direction GuardDirection) GuardDirection {
	if direction == UPWARDS {
		return RIGHT
	} else if direction == RIGHT {
		return DOWN
	} else if direction == DOWN {
		return LEFT
	} else {
		return UPWARDS
	}
}

func (UpwardsGuardDirection) getName() string {
	return "UPWARDS"
}

func (LeftGuardDirection) getName() string {
	return "LEFT"
}

func (RightGuardDirection) getName() string {
	return "RIGHT"
}

func (DownwardsGuardDirection) getName() string {
	return "DOWNWARDS"
}

func main() {
	grid := common.InputTo2dArray(common.GetInput("6"))
	var initialGuardPos = []int{0, 0}
	initialPosFound := false
	for y := range grid {
		if initialPosFound {
			break
		}
		for x, i := range grid[y] {
			if i == "^" {
				initialGuardPos = []int{x, y}
				initialPosFound = true
				break
			}
		}
	}
	guardPos := []int{0, 0}
	copy(guardPos, initialGuardPos)
	guardInMap := true
	distinctSquares := 0
	var guardDirection GuardDirection

	guardDirection = UPWARDS

	for guardInMap {
		println("Checking pos ", guardPos[0], ":", guardPos[1])
		if grid[guardPos[1]][guardPos[0]] != "X" {
			grid[guardPos[1]][guardPos[0]] = "X"
			distinctSquares++
		}

		newPos := guardDirection.transformGuardPos(guardPos[0], guardPos[1])

		if newPos[0] >= len(grid[newPos[1]]) || newPos[0] < 0 || newPos[1] >= len(grid) || newPos[1] < 0 {
			guardInMap = false
			break
		}

		if grid[newPos[1]][newPos[0]] == "#" {
			guardDirection = nextGuardDirection(guardDirection)
		} else {
			guardPos = newPos
		}

	}
	println(distinctSquares)
	loopablePositions := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == "." {
				println("Placing obstacle at position ", j, ":", i)
				newGrid := copy2DSlice(grid)
				newGrid[i][j] = "#"
				copy(guardPos, initialGuardPos)
				if detectLoop(newGrid, guardPos) {
					loopablePositions++
				}
			}
		}
	}
	println(loopablePositions)
}

type PreviousPosition struct {
	x         int
	y         int
	direction GuardDirection
}

func detectLoop(grid [][]string, guardPos []int) bool {
	guardInMap := true
	var guardDirection GuardDirection

	guardDirection = UPWARDS
	var previousPositions []PreviousPosition
	for guardInMap {
		currentPos := PreviousPosition{guardPos[0], guardPos[1], guardDirection}
		for index, position := range previousPositions {
			if index < 2 {
				continue
			}
			if position.x == currentPos.x && position.y == currentPos.y {
				//previousPosition := previousPositions[len(previousPositions)-2]
				if position.direction.getName() == currentPos.direction.getName() {
					println("Match Found")
					return true
				}
			}
		}
		//
		nextPosValid := false
		var nextPost []int
		for !nextPosValid {
			newPos := guardDirection.transformGuardPos(guardPos[0], guardPos[1])

			if newPos[0] < 0 || newPos[1] < 0 || newPos[1] >= len(grid) || newPos[0] >= len(grid[newPos[1]]) {
				guardInMap = false
				break
			}
			if grid[newPos[1]][newPos[0]] == "#" {
				guardDirection = nextGuardDirection(guardDirection)
			} else {
				nextPosValid = true
				nextPost = newPos
			}
		}
		guardPos = nextPost
		previousPositions = append(previousPositions, currentPos)
	}
	return false
}

func copy2DSlice(original [][]string) [][]string {
	// Create a new slice with the same length as the original
	sliceCopy := make([][]string, len(original))

	for i := range original {
		// Allocate a new slice for each inner slice
		sliceCopy[i] = make([]string, len(original[i]))
		// Copy the contents of the inner slice
		copy(sliceCopy[i], original[i])
	}

	return sliceCopy
}
