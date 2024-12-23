package main

import (
	"advent24/common"
)

func checkAllDirectionsPart1(grid [][]string, x int, y int) int {
	count := 0
	if x+3 < len(grid[y]) {
		if grid[y][x+1] == "M" && grid[y][x+2] == "A" && grid[y][x+3] == "S" {
			count += 1
		}
	}
	if y+3 < len(grid) {
		if grid[y+1][x] == "M" && grid[y+2][x] == "A" && grid[y+3][x] == "S" {
			count += 1
		}
	}
	if x+3 < len(grid[y]) && y+3 < len(grid) {
		if grid[y+1][x+1] == "M" && grid[y+2][x+2] == "A" && grid[y+3][x+3] == "S" {
			count += 1
		}
	}
	if x >= 3 {
		if grid[y][x-1] == "M" && grid[y][x-2] == "A" && grid[y][x-3] == "S" {
			count += 1
		}
	}
	if y >= 3 {
		if grid[y-1][x] == "M" && grid[y-2][x] == "A" && grid[y-3][x] == "S" {
			count += 1
		}
	}
	if x >= 3 && y >= 3 {
		if grid[y-1][x-1] == "M" && grid[y-2][x-2] == "A" && grid[y-3][x-3] == "S" {
			count += 1
		}
	}
	if x+3 < len(grid[y]) && y >= 3 {
		if grid[y-1][x+1] == "M" && grid[y-2][x+2] == "A" && grid[y-3][x+3] == "S" {
			count += 1
		}
	}
	if x >= 3 && y+3 < len(grid) {
		if grid[y+1][x-1] == "M" && grid[y+2][x-2] == "A" && grid[y+3][x-3] == "S" {
			count += 1
		}
	}
	return count
}

func checkAllDirectionsPart2(grid [][]string, x int, y int) bool {
	println()
	println(grid[y-1][x-1], " ", grid[y-1][x+1])
	println(" ", grid[y][x])
	println(grid[y+1][x-1], " ", grid[y+1][x+1])

	if (grid[y-1][x-1] == "M" || grid[y-1][x-1] == "S") && grid[y-1][x+1] == grid[y-1][x-1] { //if top left is M or S, and top right = top left
		if grid[y+1][x-1] != grid[y-1][x-1] && grid[y+1][x-1] == grid[y+1][x+1] && (grid[y+1][x+1] == "M" || grid[y+1][x+1] == "S") { //if bottom left is M or S, but not the same as top left, and bottom right = bottom left
			println("match")
			return true
		}
	}
	if (grid[y-1][x-1] == "M" || grid[y-1][x-1] == "S") && grid[y+1][x-1] == grid[y-1][x-1] { //if top left is M or S, and bottom left = top left
		if grid[y-1][x+1] != grid[y-1][x-1] && grid[y-1][x+1] == grid[y+1][x+1] && (grid[y+1][x+1] == "M" || grid[y+1][x+1] == "S") { //if top right is M or S, but not the same as top left, and top right = bottom left
			println("match")
			return true
		}
	}
	println("no match")
	return false
}

func main() {
	grid := common.InputTo2dArray(common.GetInput("4"))
	count := 0
	for y, row := range grid {
		for x, col := range row {
			if col == "X" {
				count += checkAllDirectionsPart1(grid, x, y)
			}
		}
	}
	println(count)
	count = 0
	for y := 1; y < len(grid)-1; y++ {
		for x := 1; x < len(grid[y])-1; x++ {
			if grid[y][x] == "A" {
				if checkAllDirectionsPart2(grid, x, y) {
					count += 1
				}
			}
		}
	}
	println(count)
}
