package main

import (
	"advent24/common"
	"math"
	"regexp"
	"strings"
)

var freqRegex, _ = regexp.Compile(`[A-Za-z0-9]`)

func main() {
	input := common.GetInput("8")
	inputLines := strings.Split(input, "\n")
	var grid [][]string
	for _, line := range inputLines {
		grid = append(grid, strings.Split(line, ""))
	}
	antennaCoords := make(map[string][][2]int)
	for y := range grid {
		for x, cell := range grid[y] {
			if freqRegex.MatchString(cell) {
				antennaCoordsList, ok := antennaCoords[cell]
				if !ok {
					antennaCoordsList = [][2]int{}
				}
				antennaCoords[cell] = append(antennaCoordsList, [2]int{x, y})
			}
		}
	}
	var antinodes [][2]int
	for _, coords := range antennaCoords {
		if len(coords) == 1 {
			continue
		}
		for i := range coords {
			for j := i + 1; j < len(coords); j++ {
				antinodeCoords := generateAntinodeCoords(coords[i], coords[j], len(grid[0]), len(grid), true)
				for _, coords := range antinodeCoords {
					if grid[coords[1]][coords[0]] != "#" {
						antinodes = append(antinodes, coords)
						grid[coords[1]][coords[0]] = "#"
					}
				}
			}
		}

	}
	println(len(antinodes))
	for _, row := range grid {
		println(strings.Join(row, ""))
	}
}

func generateAntinodeCoords(a [2]int, b [2]int, length int, height int, includeResonantHarmonics bool) [][2]int {
	dx := int(math.Abs(float64(a[0] - b[0])))
	dy := int(math.Abs(float64(a[1] - b[1])))
	if includeResonantHarmonics {
		dx, dy = simplifyRatio(dx, dy)
	}
	dCoords := [2][2]int{{}, {}}
	if a[0] > b[0] {
		dCoords[0][0] = dx
		dCoords[1][0] = -dx
	} else {
		dCoords[0][0] = -dx
		dCoords[1][0] = dx
	}
	if a[1] > b[1] {
		dCoords[0][1] = dy
		dCoords[1][1] = -dy
	} else {
		dCoords[0][1] = -dy
		dCoords[1][1] = dy
	}
	aAntinodes := generateAntinodes(a, includeResonantHarmonics, dCoords[0], length, height)
	bAntinodes := generateAntinodes(b, includeResonantHarmonics, dCoords[1], length, height)
	return append(aAntinodes, bAntinodes...)
}

func generateAntinodes(coords [2]int, includeResonantHarmonics bool, dCoords [2]int, length int, height int) [][2]int {
	var antinodes [][2]int
	if includeResonantHarmonics {
		antinodes = append(antinodes, coords)
	}
	antinode, valid := generateAntinode(coords, dCoords, length, height)
	if valid {
		antinodes = append(antinodes, antinode)
	}
	if includeResonantHarmonics && valid {
		for valid {
			antinode, valid = generateAntinode(antinodes[len(antinodes)-1], dCoords, length, height)
			if valid {
				antinodes = append(antinodes, antinode)
			}
		}
	}
	return antinodes
}

func generateAntinode(startingCoords [2]int, dCoords [2]int, length int, height int) ([2]int, bool) {
	aAntinode := [2]int{startingCoords[0] + dCoords[0], startingCoords[1] + dCoords[1]}
	if aAntinode[0] >= 0 && aAntinode[0] < length && aAntinode[1] >= 0 && aAntinode[1] < height {
		return aAntinode, true
	}
	return [2]int{}, false
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// simplifyRatio simplifies the given numerator and denominator.
func simplifyRatio(numerator, denominator int) (int, int) {
	g := gcd(numerator, denominator) // Find the GCD
	return numerator / g, denominator / g
}
