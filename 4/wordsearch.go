package main

import "strings"

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
	input := "XMMAMXMMSMXMMXMSMAXMAXSXMXXMXSMXXASMXSXMXMXMXMXXXXMASMSMAMMMSSSMASXMXSASXMXAAXXMMXAMXSXMASXAXMAMXXXMSAMSAMXXXMAXSAMXSAMMXXXSSMSMSSXMSMAAXXMX\nSXSAXXXXXXAMAMXAXXMMSMSAMXAMAMMMMXMXAMXMAMASAMMSXMASXAAMAXMAAAAXAMXMMMASMMMMMMSMMXAMASASMSMSMSMSMSASAAMSAMXMASAAMSXAMAMXSMMXAAXAAAMSASMXMASX\nSXSMSMXAMSSMMASXSSMMAAXAMXAMMXAAMASMMSAXMXAXAXAXMASXMSMSASMMMXMMMSAMAMXMAXMAAAMAMSXMAXMMAMAAXAAAXMMMMMXSAMSSMMXSXAMMSMMMXAASMMMMMMMSASXAAXXA\nSAXAAAMAAAXAXMAMMMAMMSXMMSSMSMSMSASAAMXMSMSMSMSMSSMXAAXMASMASAXAAMXSASMXMASXSMMAMMAMASMMSMSMSMSMMMSSXMASAMAMXMAXMMMMAAAXMSMXXAXMAXXMAMXXMAXX\nMMMSMSAXXASMMSASMSAMXXAAMAMASMMXMASMMSSMSAMAMAXASAMXSXXMXMXASMSMSSMSXXMAMXMAMMSMSXXMMSAAXMAXMAMXAMASAMXSSMXSAMXXAXMSXSSXXMAXSMSSXXXMSMSMSSSS\nSAAXMAXSMXMXMSAMXSASASMMMMMAMASMMAMAXSXAMMMMMAMMMXMAMMSMMMMMMAXMXAMXAMSMMAMAMAMASMXXXMMMMSMMSXSXSMASXMXMAXXSASMMXSMMAMAXSMMMSAXMMSXAXAAAXAAM\nSMSSMXMXMMAMXMAMAXXMAXSAMXMASXMAMASMMSMMMAAXMAMXMASMSAAAAXASXSSMSSMMXMAAXSSMSMMAMAMASXSAXAXXAASAMMMMMSMSMMMXAAAMASAMMMAMXAXAMMMAASMMMSMSMMMM\nXXXMXAMAXAAMAMMXSXXMAMSSMMSXXAMXMMSXAXAXASXSMSAMXXXAMXSSMSASMMAAAAXAXMXSAAAXSAMXSAMXMASASMSMMSMAMSXAAMAAAAXXXXAMXSAMXMMSSMMSSXMMMMAAAXMMXMAS\nSMMMMMSMSXXXXSAAMAMMSMXXXAXMSMSAMXMMMMSMXXMMXMASMMMMMAXAXMXMASXMMSMMMSAMXMAMMAMASAMXMAMAMXAXXAMMMMXMMXXSSMSMMSMMXSAMXMAMXMAMAXMXXXXMXXAMXSAS\nMAAAAMAMXAXSMMMXSAMAMMMSMXSAAASASAMASAMXAMAXMAXAMSAASAMXXMASMMASMXMXAAAMXXAAXAMMSXMSMMXAMXSXMASAASMSAAMAMMAMAAAAAMAXMMMXMMSSSMMMMSXAMMSMMMAS\nXSSXXXXSMSMAAXAMSAMMSASXAXMMMAMASXXAMSSMSSSMSMSAMMSMSASXMSASASAMXAMMXSSMXMMSSSSXMAXMAMSXSAXAMASMXMAMMSAAMSAMMXMMMSAMSMMAAXAAMMSAAAMXMAMASMMM\nXMAXSSMSAMSXMMSXSXMASXSMSMMMXXMASMSSMXSAAAAAAMMAMAMMSASMMMASAMXMMASAAMXMSAAMXAXMMMSMAMAMMAXXMAXMSMMMAXXMMSASMMSAXAMSAASMSMMSMAMMMMSAMMSAMAXS\nMMAMMAAMAMAXSAMXXXXXMMSMMAAAAXSAMAAAAAMMMXMMMXXAMXSAMXMAAMMMMMMSSSXMASAMXSASMSMSAAAXAMAAMMMSSXSAMSXMMMSSMSAMAASXMASXMXMXAXAXMXMSXMSASAMXSXMM\nASXMXAMMXMXXMASMSXSASAMASXMMSXXAMMMSMXMSSSMSXSSXSAMXSMSXMMMAXMAXMAXXAMAXAMXSAAASMSMSMSASAMXAAASMXAXXXAAAAMMMMMSXAXAXXMMSMMMSMSXSAASAMXSMMMSS\nMMASMSMMXSSMSAMXAASAMSMMMMAXXASXMMXMASXXMAAMAMAMMXXAMXAXMASXSMSSSMSMMSAMSSMMXMMMMXXXAXAMASMMSMXXSMSMSSSMMMAMAMXXMMMASAMXAAAAMSAXMMMXMMMXAAAA\nXMAMXXAXAMMXMASMSMMXMAMSASAMXXAXASXMAMMASMMMAMXMXMMSSMMMSASAAAAAAXAAAMAMXAMAAMASXXMMSMXSAMXAMXXMAXAAXAAMSMXSASASXXMASMMSMMSSSMMMSAMXSXMSMMSS\nSMSSMSSMMSAXSAMMMAXXXSXSASMMMSMSMMAMMMMAMAASMSMSXAMAAAAAMXMMMMMXMMMSMSSMSAMSXSAMAMXAAAXMASMSSMMXMXMXMSMMAAMSAMASXXMASXXAXXAMAMXAXAXAXAMXMAMX\nAMAAAMXAMXAMMAMXSXMMMMAXXMAXMAAMASXMASMMSMMXAAAXAMMSSMMMSAMXXAMAAAXAXXAAMXMAMMASXAMMSSXSAMAMAMSSSXMSAXXSMMMMMMAMAMMXSMSSXSAXAMMMSMMMSAMXMASM\nMMSMMSMSMSXXMXMXAMMAAMMMMSXSSMSMAMASASAAAASMSMSSSMAMAXAXSASMSASXSSSMSSMMMSMMXSAMMSSMMAMAAMAMAAAAMAAMSMMMAAMAMMXSSSMSMMAMMSAMSSMMAMAAMAXSSMSX\nXAXXMAAAAMASMAMXMASXSSMAXMAXAAXMMMXMASMMXSMAAAXAAMXMMMAMSXMXSASAMXAMXMASXXAAAMASAAAXMASXMSMMSMMMMMMMAMASMXSASXXXXAXSAMASXMAMAAMSMSMSSMMMAAXM\nMMMMMMXMAMAAMASAMMMAMAXSXMMMMAMAXXMMMMXAMMMSMSMMMMSXMASMXAMASXMASXMMAMMMASMMSSMMMSXMXAMAASAAXASMXMAXMMAXAASMMMSSSMASAMASXSSMSXMAXXAAXXASMMMX\nAMASXMMSAMMSMMMASXMSSSMMXXAAXAMXMMSAMMMSXSAMXXAXAXMAMXMASXMAMAMAMMASXSASMXMXXAAXXAASXMMSMMMMMAMXAXMSXMMSMXMMAXAAAXMSAMAXAXAAXSSXSMSMMSMSXMAM\nXMASAAASASXMAMXAMXAAAMAMMSSSSMSMSAXASMAXXMMMMMMXMMSAMAXMXXMASMMSSXMXAMAMXMXXXSMMMMMMAXAMXMSXMSMSMSAAXAASMMASXSMSMMXXAMMXSXMMMMXAAAAAAAAMAXSA\nSMASXMMSAMXXMMASMMMMSMAMAXMAXXAAXSSXMMASMMAAAAMASMSASMSXMSMASAAXXMXMSMXMASMMMXAAMASMMMXSAMMSMAAAAAXMXMXXASAMXMXAMSMSSMSAAASMSSMSMXMSSSXSAMXS\nXMAXAXMMMMMMXMMMAXMAXMASXXSSSSMSMASXAMASASXSSMXAMAXAMAAXAAMXSMMMSMAAAAMSAMAAXSSMSASAXAASAMAAXSSMXMAXSMMMAMXSAXSASAAMMAMMXMMAAXAAMMMMAAAMMSXA\nSSXSAMAAAAASXMXSMMSSMSXSXMAMAAAMMAMSMMASMMAMAAMSMMMSSSMMSASASMAAAXSSMSMMASXXMAXXMASMMMMSSMSSMMAXXSMAAXMASXXMAXSXSMMMMMXMXSMMMMSMMAAMMMMMASXM\nMAMMXSMSMSMSAMXAAAAAXSAMXAMMSMMMMXMAMSMMMMSMAMAAAMAMAAXXMSMASMXXMMMAMAMSMMASMMMMSAMXSAMMAMXMAXXMMASXSSMXMASMXMMASAMSAMSAMXAAXXMASXSSXAXMASMX\nXMASAXXMXXXSAMXMMMSMMXMAMXSAMXSMAMMAXSAMXAXXXXXSSMAXSMMXXXMAMMSMXASAMMMMMMAMAAMXMAMXSAXXAMMSSMSASAMXMXMXMASXMAMAMAMSASXMASMMMAMAMAAAXMXMASAX\nSXMMXSAMXMMSAMSXSAMAMMXMSMAAAAMMASMSMSAMMMXMMSAMXXSXMXMMXMXSAAAAMMSAMMAAAMAXMXMSSSMASXMSXXAAAASMMXSXMASXMXSASXMAMSMMMMASMMSSMXMAMMMMMMAMAMXM\nXXMAXMAMMMAMAMXAMAMAMAAMAASXMASMXSXAASXMXSAAAXAXAMMASAMXSMAXMSSSSMMMSSSSSSXSMSMMAAMXMSXAXMSSMMMXXMAMSAMXXASXMAXAMXAMXMASMAAXAAXMMSMSAXXMMSSM\nAMMXXMAMXMSMMMMSMMSSMSXSAXMAMXAMXXXMMMXSASMSMSSMASXSMAMMSMASXXMAXXAAAXXAXXAAAXAMSMMAAMMMSXAAXMXMASMMXSXXMASMSSMMMSAMAMASMMSSSSXSAAAMMSMAAAMA\nSXSAMSAMMXMAAAAMAXAMMXMAXXMXSAAMXMASXMAMMSAMAAXSAMXAXSMAXMASAAMXMAMMSSMMMMSMMMSMAMMMAMXXMMSSMSMSASMSAMSXSAMMAAAXXSAMXMAXAXMAMAAMSMMMASMMMSSX\nMAMAXXMMXAXXXMSMSMMSSXMASMSAMXSMSAXXAMASXSAMMMXMMSMMMXXSXMXXXAMXAXXMAMXAAXXAAAMMMMSSMMSAMXAMAAAMAXAMAMXXAAMMMSMMXMASMMMSSMSAMMXMAXAMXSAXMXMM\nMXMSMSMSMMSXSAMXMAXAMMXMXXMAMAMAXSAMXSXSASXMASAAAAMXMXAMMSMMXMAXSAAMXSSMSSSMMMSAXAMAAAMAXMASMMSMSMSMSAMXSMMSMMMXAXAMAXAAAXXASXMSSSMSXSMMXAXA\nSXMXAAXMAMAMMAMAXAMXSASXMSXSASMMMXXMAMAMXMASXMSSMSAMXMSMAAAXXXMAMMMXMXMXAMAAAASAMSMSMMXMMSMMXSAAXAAAMAXAXAXSAMSSMSSSMMMMXSSMMAXAMXAAMMMSSMSS\nMAMMSMXSAMSMSAMXSMAAMMSAAXAXMXAXAMSAASAMMSAMMMMXMXMAMMAXSMSMASMMMXMAAMMMSSSMMXSXAXAXASASAMXMXMASMXMMSSMMSSMSAMAAAAAAXMASAMASMXMAMMXMXXAMAAAA\nSAMAASXMXMXASAXMAXAMXAXMMMMMXSXMSAASXMASAMASMAXMASXSSMMMMXAAXAAASMSMXSAAMXMASMMMXMAMASASXXAAAAXMXXXAAXAAXMASMMSMMMSMMMSMMSAMXAXMSSMMMMMSMMMM\nSASMMMMSSMMMMMSMXSXMMXXAAAAAAAASMMAMSMXMMSXMXMSMMXXAAMXAMSXSAXSMSAAAAMMXXASAMAASXMAMXMAMASXSMSMSSSMMSSMMSMAMMAXAAMAMAXMAMXAASMMMAMAAMAMXMMSM\nSXMAXSAAXAXXAXAAAXAXAXSSSSSMSSXXAXXSXSAMXMMSXMAXMASAMXSSXMAXAMMAMMMMMSAMSAMXXMMMSAMXAMXMMXMAXAAXXXAAXAAAXMAXMXSSMSASMMSAMSXMAAAMASMMSASMMAAX\nSAMMMMMSSSMSMSMMSSMMXMAMXAAXAXMSMMSMASMXAAMAASAXMAMXXAXAAMAMAMMAMXXMXMAXMXMXXSXMXMAMXSXXXXAMSMMMMSMMSMMMSSXSAAXAXMMMXASAMSXMXSMMMSAXSAMAMSMS\nSXSXMSAAAAAAASXXXAMASMSMMSMMMSXAAMAMMMSSXMMSAMASMSSMSMMMXMSSMMSSMAXSAMSMXAASMMMSAMXMXMMSASAMXXAAMXAAXXAXAAAMMMSAMXAXMXXAAMASAMASXMXMMAMXMXAS\nMAAXMMMSMMSMSMMMMAMASAAMMAXAXAXSSSMMSAMMSXAXSMAMAAMAXXMSMAAAAXAXMAMSASAMXMSAMAAXASXSMAAMAMAMAXMSASMMSSMMMSMSXAXXASMSSSSSMSAMASAAMMAMXAMSAMAM\nSMMSSMXXAXXMAAMSSMMMMXMSSMSSSSXMAAXAMAMAMMMMXMAMMMMSMMMAAMSMMMMSMSXSXMASXSMMSSXXMMAASMMSMSMMXSAMMMAMMXMAXXXAMASMMAXAAMAMAMXSXMAXSSXSAASMXMAM\nAAXAAMAMAMMSXSAAAXXAXXXXAMSAAMMMSMMSXMMSMSMAXASAMXSASASMMXMAXSXSAMAXMMXMMAAXAAMSSMSMMMAAAXMASMAMXXAXMASMSMXMASXMSMMMSMMMMMAXMASXMAMXMSXMASXS\nMMMSSMASMMMAMMMMSMSMSMMSAMMMMMAXXXAXASAXAAAXAMXAMXSASXSXXSSSMAAMAMMMASMMSSMMMXMAAMXMAMXSSXMXSMMMMSMXSAMAAAAXSXMXAXAXAXMASMMXSAAMMAMMSMAMMSXM\nXMAXAXMSASMAMSSMXAAAAAXXAMXAAXXAMXXSMMAMSMXSSXSXMMMMMAMMXAAAMMSMAMAAMAAMAXXMAXMSMMASMSAMMXMMMMMAXAAXMAXSASXSMAMSMSMSMMSASAAMXMXAMAMXASAMAMAS\nXMXXAMXMAXXAMASMMSMMSSMSSMSSSSMXXSXMXMAMAMXAMAMXMXMMXSASAMMMSAXXXXXSXMSMASMMSAMMAMASAMXSXAXAAAMSSMSMMMMXAAXAXXMAXAASXXMSSMMMAXSSXSSSMSXMASAM\nMSSSMSXMSMMXMXSXAMAAAAAXAAXXMAMSAMMSASASASXMMSMSMSMAAXSAMXMAMXSMSSMMMXMMMSAAXAASAMXXMAMMMXXMSXSAMXMMMSAMXMXMMMSXSMSMXXMASMASMMMAMAAMXMASAMAS\nAAAAAAXAAAASMMMMSMMMXMMSMMMSMAMMAMASASASASASAMXAAAMMMSAMXAMASMMAAXAASASAAMMMSMMSAMXSMSASXSSXAMMXMAMAXMMSASXSAAAXAMAMMXMXXMAAMAXAAMXMASMMASAM\nSMSMMMMASMMMAAMXMXSMXSAMXAAAMASMMMXSXMAMXSAMXMSMSMSAMXAXSXSASAMAMSSMSASMMMAXXXMXMMMAAMASAAAMMSXXXXMSMMXSXSASMSSSMSAXMXSMSMSSSSSMSXASAMAMAMAX\nXMMMMXMAMAXSMMSSMAAAAMASXMXMSAMAMXMMMMMXXMAXMASAXXSMSSXMSAMXSAMMMAAXMAMASXMSMSSMMAMMXMAMMMMSASXXSMAXAMASXMMMXXAXXMMMXAAXMAAAXMAXMXXSASMMSSSS\nXSAMXAMXSMMSMSAAMSMMMSAMASXAMASAMSAMXAAMXSSMXXMAMMMSMAMMXAXASAMXMSXMMAMAMAMAMAAASXSXAMXMAMXMASMMMSASAMAXASXXXMAMMMSAMMSMMMMSXXXMXXMMAMAAAAAX\nSSSSSXMXAMAXXMXMMMMSAMXSAMMMMXAMXXASXMSXMMMMMSMMSMSAMXAXXAAXMMSAXMMSSSSSSSMAMSSMMASXXSASASAMMMAAASMSXMASXMMXSMAMXAMXMAAAMAMXMSMMMMAMXMMXMMMM\nXMAXAMXXAMXSMXSSMSXSMMMMASAXXXMASMMMAXMASXMAAXAAAXMAMSSMMSMXAAMXXAAXAAAXAXXXXXAAMSMMMSASXMXMASMMMSAMXAMXMASMMMAXMXSASMSMMSSXAAXMASMMMXMAXAXX\nAASXMXAXSMMXAAXAAAXSASXSXSMXXAMXMAMSMMMAMAMMXSMMSMSAMMAAAXXSMMSMSMMMMMMMMMXSMSMMMMAAAMMMXSASAMAMAMAMAXXASAMAASMSMXMAXXAAAXAXSMXXAMAMAASXSMMS\nXXXMXMMXAAMSSXSMMMXSAMASAMAMSSMXMAMSASMASMMSAMXSAMXMASMMMSXMXMAAAAASMMAAXAAMAXSAMXSMXSASXSAMMXAMAXAMAASXSASMMMMAMXMMMSSMMXMMMXAMXSSMSMSMAMAX\nMMMAXAXXSAMXXXMASXAMAMAMMMAMMXAAXXXSAMSAMMASAMXMAMXMASMAAXAAASMSSSMSASMSXMXSSMXXSAXAXXAXAMAMSSSXMSSMXMSASAMXSXSAMAAAAXMXSAXAXMMSAAXAAAMXAMSS\nAAAMMSMMMASAXMMMAMXSXMASXSSSMMMMSMAMXMMMXMXSMMASAMSMMSXMMSMSMSAAAMXSMMXMASAAAMAMMMSSSMMMSSMMMAMAMAAMXAMMMAMAAASMSXSMSXSASMMSXMASMMMSMSMMXSMX\nMMSXAAAXSAMXMMAMAXAMXMXAMMMMMXAAAMXMMSAMXSAMAMASASAAXMMSAXXXAMMMMMAMXSASAMMSMMAMAMXAMXMAXAAXMAMMMSSSMMSXSXMMMMMAMAMAAAMASAMXAMASAXAAXXASMSAA\nSAXMSSXMMMMAAMAXSMMAAXMMAXAAMMMXMMAMXSASMMAMXMASMSMXMMAXXAMMMMXMAMXSAMXMASAXXSXMAMXMSAMXMXSMMMSMAMXXAXMASXXAXASAMAMMMSMAMAMXAMASMMSXMXAMAMXS\nAAXAAMXSAAMMSSMSAASXMSSXSSMSMSAMSMMXMXMAAMMSSMXSAMXSXMMMMAMAMSASASXMXSAMAMXMAXMASMSXSASMSMMAAAAMAMSSSMMAMXSASXSASASAAXMMSSMSMMASXMASXMSMSMAX\nMSMMMSASMXXAAAXMMMMAAAXMAXXAAMSMAASXMAMMSXAAAXAMXMASXAAMSASAXMAXAXXASMXMMSSMSMMXAAXAXAMAAASMMSSSXSAMXAMXMMSASXSMMXXMSSSXXAAAXMASAXAXXAAAAMXM\nSXASAMASASMMMXMXAMSMMMSMAMMASAMMSMMASAMXMMMSSMSSSMAMSSMXSXSMSMSMMMMMMMXAXAMXXAMSMSMSMSMSMMMSAAXXMMMSSMMSXAMAMAXMMXSXXAMXSMMMXXAMMAMXMXMXAMXA\nASAMXMMMAMMAMMMXMXXMAXAMAXSAXASAMASMSASAXAXAAMXAAMMMXXMASMSAAAAXMXAXAXMAMXSASAMAAAAAMAAAXAASMMSSMAMMAAAXASMSMSMASMASMAMMMAAXSMAMSSMASXSMSAXS\nMMMMXSAMSMSASAAXSASXSSMSAXMXSAMASAMXSAMASMMMSMMMMSSSMAMAXXMSMSMMMSMSMMMSXXMASMMMSMSSMMSMSMMXASAMXASMSMMSAXMAAMSAMAXXMAMAMMAASXSMAXXXSAMAMXMX\nXMAMAMXSAASASMSMAAXMMAXMXSAMXMMXMXMAMAMXMAAAAAXXXXAASMMMSMMXXAXAXAAAAAAAAXMMMAAXXMAXXMAMXAMXXMASXMSAXMASASMMXMMXSSMMSSSSSXMXXAXMXXMXMAMXMSAM\nMAAMMMMMMXMAMMXXMSMSSXMSAMXAAMSXAAMAXAMXSSMSSSMSMMSMMSMAAAMAXMSSSMSMSMSSSMMAMMMXSMSXXMAMSAMSSMAMAXMXMMAMXMAXSAAAMAMXAMMAMXMMMXMSAMSASAMAMSAM\nAXSMSAXMMMMSMMXAXXAAASXXMASMMXAMSMXMSMSMXXAMAAAXXAAAXXMXXMMXSXAXMAMAXMXAXMXAMXMASXAAMSAMSAXAAMSSXMSAMMMXSSXMASMMSMMMMXMASAAASASMAXSAMAMXMSAM\nSMMASASMSAAAAMSMMMMMMMMAMXMSSXSMXXAMAXAASXXMSMMMMSMSMXXMSSMMAAMXXMMMAMXMSXSXSAMXSMMMMAMMSAMSSMMXMASASXSAXMAXXAMXAAMMAMSAMMSASXSMXMMMMSSMMSAM\nMAMAMAMASMSSXMAMAXXASASAXAAAMAMAMSMSAXMSMAAXXSAAXXAXXXMMAAASXXSAMXASAMMXXMAASAMXSASXMAMMXMXAMAXXMAMAAAMSMMSMAMSSSXSMAXMAXMMMMAMMAMXMAAAAMMXS\nMAMMMXMXMXXMASASMSSMMASASMSMMAMAMAMAMASAXSXMASXSMMAMASMMSSMMMAAMSSMSASAMXAMXMASXMXMASXSXAXMASXMXMXSMMXMAXMASAAMAXAMXSMMSMXAAAAMSXSAMMSSMMMAS\nSMSMMASAMMXMMSASXMASMAMAMAAXMASASMSAMXXAXMSAXMAMAMMMXMAAAXMAMXMAAXXXMMMXMSXMAXMXMAMMMXAXXMSAMXSSMMXXSXSMMSASMSMAMMMAMAAXMSMMMMAXXSXMXAMXAMAS\nMAAAXXSAXXAMXMAMXMAMMXMAMXMMMAMAXXMASXMSMXASXXASMMXMXSMMMSMMSMXMASMMSAXMSAMXMSSSSSMXSMSMSAMXSXAAAXXXMAMAMMMMMAMXSAMSSMMSAASXXSASMSASASXSXMAS\nMSSSMMSAMSXMXSMMSMMSAMXSMAXSXSMSMSMSAASASXAXXMASXSAMXXAMASAMXMAXMAMASAMSAMXSXAAMAMMXMAMAXXSXSMMXMMXMMSMAMAAMSMSXMMSXAXAMMMXSAXASAMAMMMAMSMXS\nXXAXMAMAMMASAAAAMAMXAMAASXSAMXAMAMXAMMSAMMXMXXAMXSMSMSXMASMMASMSMAMMMXMMAMAMMMSMAMSAMMMMMAMAXSMSMSXAMASASMSMMXMAMXMXMMMXSMXMMMMMMMXMXMAMAMAM\nSMMMASMSMSAMXSMMSAMSMMMXAMMAMMXMMMXMXXMMMMASXSMSAMMXMXMMMXAMXXAAMASXMAXSAMXSAMXMMSSXSASAMAMAMXXAAXMMSMSXSAXXMASAMASAMXSAAMXMAAASAMMMXSXSAXAM\nAAAAAMAXAMXSAXMAMXXXMAMSMXSAMMMMAMMXMSASXSASAAAMMSMSSMMAASMSSMSMSMSASAXSMSMSXXSAMXMASASXXAMMSMSMSMAAMXMMMXMASXSASASXSAMSMMAMASXSASMMAMMSASXS\nMMMSMMMMAMAMMXMSMXAAMAMXXXXAXAAXAXMAMMAMXMXSMMMMAAAMMASMXSAAXAXXAASAMSMMAAASMMSASAMXMAMMMMXAAXAAAMXMMAXMASXXMASMMMSAMXMAXSMSAMMSXMAXAXASAMXA\nSSMMXAAMAMSXSAAASXMMSMSMSMSMSXXMASMXMMSMSSMSXXAMSMSMSXMMMMMMMMMMMMMAMXXMMMSMAAXMAMSAMXMXASMSSSSSSSMASXXMASMXMMMAAXMXMXMAMAAMMMASAMAXMXAMAXMM\nXAASXSSXXMXASMSMSMXMAMAMXAAAMAXSASAMXXMASAMXAMXMAAXAMAAXSAAXXAAMXMSMMMXMASAXMMSMMAMXXAXMXSAXAAAAAAMAMMAXAMXXMXSMMSSMMAMXSMXMASMSAMMXSASXSMAM\nSMMMAAAASXMAMXMXXAAMASASMSMSMAXMXMXAMSMMMMXXMMASMSMAMSSMSSSMSMXSAAXMAMXSXSAMSAAAXSMMSASASMXAMXMMMSMSXSAMXAAXXAMAAAMAXASASXMSASASAMAAAXMAMXAM\nAXSMMMMMXAMMMASMSSXSASASAAAXMMMSAMXXAAAAAXAMXSXSXMAXAXXMXAMXMAAMMMMSAAAXMMAMMSSSXXAAAAMXXSAMXSXAXMAXAXXMSSMXSASMMMSSMMMAXAXMASAMXMMMSXMXMXMM\nMMMAXMXXXMMASASAAAMMMMXMXMMMAAAXAXAXSSSMSAMXAMXSAXSMMXSMMASAMMMXSAAAMMMXAMXMAXMAMXMMSXMXMAMXAASXXSAMXMMXMAMXMAMAAXAAAAMAMSMMAMXMXSAAMXMAXAMS\nXAXMMMMSXAAXMAMMMSMMASAXMMASXMMXXAXXXAXMAMSXXSASXMXMSAMXASAMSMAASMSXMASXMMSMSXSASMSMMXAXMXSMSMMXMXSAXMMXSAMXMXMSXMMSXMXSXXAMXMASASMMXAXMSMSA\nSSXXAAAMMMMSMSMAMAXSASASMSASAAASXSXSSSMAAAMMMMMSXAAXMASAMMAXSMMXSAMXMAXAMAAAMMMAXAAAMMMMMAXMAAMAMAAXMAAASASXMXAMMXMAMSAMMSMMSMMMASAMSXXXAXXM\nMXAXSSXSASASAAMXSMXMASASAMASMMMSAMAMAMXXMSMAMAASXSSSXAMAXMMMSASAMMMAMMSMMSMMSAMAMSSMAAAAMSSMSAMAMMSXSAMMSSMAMSASAXMAMMASAXAAAAXMAMAMAASMSMSX\nMMMMXAXSASAMXMMAAMAMXMAMXMXMASXMXMAMAMXXMAMXXMSXMXAXMASXMSAASXMASXSMMAAMXMAASXMXMXAXSSSMXAXMMXXAMXAASXMAMASAMMAAXSMSMMMMXXMSSSMMXXXMMSMAMAMM\nXAXSMSMMMMMMMMMSMSMSAMAMAXAMXSAMMSMMXSAXSASXXSMXSMMMSXSAAAMMSAMAMXMAMSXSAMMAMMMXMSSMAMXMMSMMMASXSMMMMSSXSAMXSMSMSMAXMAMASXMAMAMMMSMXMXMXMAMA\nMXSXAAAAXAMASMAMMXXMAXAMXSMXASAMAAAXAXXAMAXAMXAAMAAAXMMMMMSAXXMMMMMAMXAMXXXMMASAAXXXMSMXAXAAMAMAAAMAAAMXMAMXSXAXAMMMSMSAAXMASXAAAAAAXAMASASM\nMXSMSMSMSXSASMMSAXSXSMSMMMMAASAMSSSMMSSMMMMSSMMMSSMSSMAAAXMMXMSAAMSMSMSMMMSASASAMXSMXMMMAMSMSASMSMMMMMXASMMAMAMSSMSAAAMAMXSAMXXMSSMMSMSASXSM\nXASMMMMXXAMASAMMMMSAXAXXAAXXXSAMAMAXAAAXXAAMAMXMXAXXAMSSMSASAAMSXMAAAXAAAMAAMAMAMASAAXXXXXXXMASAAAXXAAAXAMMSMSXMXAMMSSSMXMMASMSMXMXMAXMASAMM\nMMSAMAXMMXMAMAMMMAMMMMMSMSSMXSAMMSMMMSSMSMSSXMAMMMMSAMXMASAMMSMXXSMSMSSSMXMSMMXSMXXSASMXSMXAMXMMSSMSMSSMASAMXXASMSMXAAAXAASMSAAMAMSSMSMMMMMA\nSAMAMMXAMMMMSAMAMXXAAAXAXAAXAXXMXAXSAAAAXAAMASASAMXMAXAAAMAMAAXMXMAXXXAMXXXXXMAXMSMMSAMXAMASMMAXAAXMAAXMAMXSMSAMAASXMAMMMXMAMMMMASAAXSAMXXXM\nMASXMXSSXMAMXMSSSMSSSSSMSMSMMSSSXMAMMSSMMMMSMSAMAMXXASMSSSSMSXSAMMSMSMAMXMASXMASASAMXMSXMSAXAXXMSSMMMMMSSXXXAMMMMMAXXAXXSAMAMAMSXMXXMMSAMXMX\nSAMXAMXAXSASMMAAAAXAAAAAXXAAXAXMAMAMMAXXMXXXAMAMSMSMAAAAXXAAAAMXMXAAAXAMSAMSAMXMXMAXSXMXXMASAMXMXAAAAAAMMMMMSMXAMXMASAMXXAXAMAXAMXSSSMXXXAXM\nMXMXMMMSMSASAMXSMSMMMSMMMSSSMMSSXMAMSAMSAMXMMMAXAAAMMMMMMSMMMSMXSMMSMXXMSAXSXSAMXSXMAAMSAMXMXSAMAMSSSMMMAAMAXMSMXMXMXAMXMXMXXSMMMMXAASAMSMSA\nXMSSMSXXASAMXAMXAAXXAXMAMAMXMXXAXMAXMASAASXMXSMSMSMXAAXMASMAMAMXMAXXAAXMSMMSAMMXMASAMXMAAXMAMMAXXXAMAASXSSSMXMASXSMXSSMMASMSAAXSAMMSMMXMAMAX\nMXAAMXAMAMXMMXXMMMXSXMXASAMASXSMXSMMMMMMXMAMAMMAMAAXSSSMSXMAMAXXMSMSSXSAXMAMAMAAXXSMXXSMSMSAAMSMSASMSMMAAAXAASAXMASAXAAXMAAMMMMXAMMAMMXSMSMS\nMMSXMMSMAMXMASXMASAMXSMMMAXAXXXMASAXAAAXXSAMXMSASMSMMAAXAMSSSMSSMAAXAAMMMMMSAMXSSMMMSMSAXASMMMAAXAMAXAMXMASMMMXMMMMMSSMMXMMMXSMMSMSAMMAMMAMS\nMAXAMAMSAMXMAXASAMASASASMXMSMMXMASASXSXMXSMSMXSXXAMXMMMMMMXAAXAASMSMMMMASXASASAXAASASAMAMAMAXXMMMAMMSAMXMAXMXXAMSXMAAAXMASXMAXMAXAMMSMASXSXS\nMASAMAMXMMXMAXXMASXMASAMXSAXASXMMSAMAMMSMMMSMXMASMMAXAAXXAMMMMSXMXAAAASXXMASAMASMMMASMMXMSMXMASXSXXAXAAXMXSSMSMXAAXMSMMXAMAMXSMAMAMAAXXMMXAM\nMMXAMSMSAMXMMSMSXMAMXMXMXMAXMMMSAMAMXSAAAAAMAMXAAASMSSMSSMAAXMXASXMXSXSXSSMXMAXMXAMXMAMMMAXAMXAASMMMSSMXSAMXAAMXSXMMAASMSSSMMMMMSXMSSXSAAMAM\nASXXMAASXMAXXAASMSMMMMAMXMSMAMXMASXMAMXSSMSXMMXXSMMMAXAMXXSXSAXMMAXAMAMXAAXXMMMMSSMAXAMSMMSMSMMSMMXAAAAAAAXMSMXXAAMSAMXAXMXAAAAAAAXXAMSMMSAM\nASAMXMXMASASAMMMAAXAAMASAAAMAMXSAMAMXSMMXXMASXMAXAAMXSMASXMXMMXSSSMXMASMXSMXSAMXAASMXSMSAXAAAMXMAXMMMSMSSSMAMAMSSMMAMXMSSMSXMSMSSMMMAXSAMXMM\nXMASXXMSMMAXMXXMSMSSXSASMSMXSSMMSSXMMSASXSSMMAAAMSSMMXMXSAMMMXMXAXAASXXMMAMXSASMXMMXAXAXMMMSXSASMMMSAMXAAMAAMAMMAXXSXSMASAAAMXMMAMXSAMXAMSMX\nMSAMXAMSXMAMXSMAMXAAMMMMXAXAXAMXAMASASMMAXAASMXXXAMXSAMAMAMAXAMMAMSMXMAMSMMMSMMXSXAASMSMMXXMXMASAAAMAXMMSMSMXAMSXMAMMXMAMMSMMAXSAMXAXMXMMASA\nXMAXMSMMAMSAMSMAMMMSMMMSSSSMSAMMMSAMMSXMMMSMMXMSMMSAMAMAXASXSSSMAXAXMMMMAAXASAMAAMSMAAXSMAAAMMMMXMSSSMSAMXAMXSXSASAMXSMSMXXASAMSASXMASAASASX\nMSSMMXASXMXMMXMASXAXAXAXAAAASAMMXMASXMXXSAXAMXAXAAMASXSXSXSMAAXXSXMXSAASXSMXSAMXMAASMSMMSSMSAAMAMXXAXAXASMXSAMAMAMASXSAXSASMMXXXXMAMAMMXMXSA\nMAMAMSMMAXAMSSSMSMMMAMXMMMMMMXMXXSXXAAAMMMMAMMMSMMSAMXAXAMXMSMSXMAXAXSXSAXMAXAMXXSXAXMAMXMAXXSSMSSMMMMMXMXAMXSAMSMXMXMAMMXMXXMASMXMMASXMMXXM\nMAMAMAASXMSAAAAASASAAAXMXXXMAXSXMMMSXMXSASMMMAAAAXXAAMMMSXMMAAXASMMSMMMMMMSSMAMSXMSMXSAMXMAMXXAAAAAMSXMASMMXAMAMXXXSXMAMXASXSASAMXSMMXAAAXSX\nSSSSMSMMXAMXMXMMSAMMSXSMSMMXSAMXAAAXMAMSASXSXMSSSMMXMASAMASXSSMMMXAXXAAASXAAMAMXSAAXAMASAMXSXSMMSSMMMASASAXAXSAMXMAAAMSSMASAAMXAMXMASXMMMXAX\nXAAMAXSMMSMMMSMMMMMMXAAMAMSAXASXSMXSASXMMMXSXMMMMAXAXMMASAMXAAMXMMMSSXSSSMSMMXAMMSMSXSAMXSXMASXAAAAXSAMXSXMMMAMXMAASXMAAMAMMMASMMMSAMAMSMMMX\nMXMMAMAMAAAXSAMMAMAAMMXXMXMASAMMXMAXMMAXMAMMASXSSXSASXSMMMSMMXMAMAMXXAMAXXXSSXASAMXAMXMSXXAMAMMSMSMMMASXMASXXXXMMSAMXMSMMASXMASXAAMASMXAASXM\nMMXMXSAMSSSMSAMSASMSSMSASXMAMAMXSXSMAMMASAXMAMAMAAMAMAASAMXXSMSXSMSAMXMXMSMAMSXMASXMMMMMXMXMAXXXAMMMXAXXSAMMMMXMAMAXXXXASXSAXSMMMXSAMXMMSMAS\nASAMXMAXAAAXMAMSASXMAAMAAXMAXAMASAMXAAMASMSMASAMXMMSMMMMMSAXXAMAMAMASAMAMAMAMXXSMMXAAXAXMASMSMAMMMAAMMSXMXXAAXAMMSMMMXSAMAMMMMAAAASMMMMSAMAM\nXSMSASMMXSMMMSMMAMXSMXMXSXXSSXMAMAMSMSMAMAAMXMMAXMXAXXXAXMAXMSMXSASMMXSMSASMSMASMASMMXAMSXXAAMAMMSXSAMSXMSSSSSMSAAAAAXMAMSMSASXMMMSAAASMMMSS\nAXASMSXMAXMXXASMXMMXXXSAMMXMAMMASMMSAAMAMSMXSAMXSMSMMXMASMXMSAMXSASAAXXASASAAMAMMAXASASXXAMSMSMXMSAMMAMMXAAMMAASMSSMMXMAMXASAXXMSASXMXSAAAXM\nSMMMXXXXSSXSXMASMAMMMXMASASMAMSAMAAMMMMXMMAMMSMAAXMXSXSSXMASXASMMMMMMXMMMMMMMMXSMXSAMAMAMXMMMMXAAMMMMAMAMMSMSMMMAMXAXSSMSMMMSMAMSAMXMASMMMSS\nMASXSSMAMMAMAXAXMAMAMASXMMSMAMMSMMXSXAMXAMAMAAMMXMASAMXXMMMXMXMAMXSASXASAMXMXSASAMMAMAMAMAMXAASMSMSXMAXMMMMAXMAMAMMMMMAAXXAAAMXMMMMXSAMASAMX\nXSMAXAMAXMSMAMMSSSMXSASXMASXMSAMXSMMXMMMMSMMXSSMSSMAMASMMASXSMSAMXSAMSMMXSMMAMMMAMMMSXSASXSSMMSXAASMSMMAAAMMMXMXMXSAASMMMSMSXXAAAMSMMXSAMXSA\nXMMMMXSMSMXAMSXMAAAMMAMAMXSMSMMMAMAMXSXMAAAMSMXMAAXXXXMASAXXAASASAMXMMMSAMAMXXXXAMAMSASASAXXXAMXMSMAAAMSMMSMSAAASASAMSAAXXAAXSSSMSASXXMAXSAM\nAMASAAMMAMAMXMAMMMMXAAMSAMXXAXXMAXSMAMAMSMSMAASMSSMSAXSMMSSMMMSAMASXSMASMMAMMXSSXSASMAMAMMMSMSSXMAMMMSMMAXMASMSMSAMXMSXMMMSMAXMAMSAMMMSASAMX\nXSAMMAXSASASAMAMXXXMMMMMASXSSSMXSXMMMMAMMSMMMXMXMAAAAMXMAXXAXXSMSXMAAMAMXSSMSAAMAMMSMSMSMMAMAAAAASAMAXASXMMXMXAXMXMAXXAXSAMXXAMSMMAMAAXMAXSS\nMMASXMASAMAMASXSSMMAXMXSMMAAAAAAMASXSSSSXAMXSAMXSMSMMAAMXXAXMMSAAXMSMMASXXAAMMSMAMSMMAAAMMAMMMSMMMAMXMAMMXSAMXSSMXSMSSSMMASXXAMAXSXSMMSAMXXM\nMSMMMSMMAMMMAMAAXAMSMSAMAMXMSMMMSAMAMAMAMASAMXMXMXXMSMMSMMSMSAMAMXMAXSAMXMMMMAXMMMAAXMSMXSAXAXXMXMXMAMSXMASASAMXMASAAMAASXMMSSSSMSAMXMSXXMSX\nXAMAAAXXXASASMMMMXMAAMAMMMMMAASXMAMXMMMAMXMASMMAMMMXAMSAMXAAMMMAXAXMMMASAXXASXSASXSSMXAMXSASXSSMSSMMXSMAMASAMXSAMMMMMSSMMSXAXXAAMXMASXMXAAMX\nSASMSSSXMAAAXAAXXASXSSSMAAASMSMASMMMSASXSXMXXASMSAMMSXSASMMSMMSSSMSXXXMAXMXASAMAMAXAMXAMAMAMXAAXAAAMMAXXMASXMAMXSSXSXAXAXSMSSMSMMSXASAMXMASM\nAXXAMAMMMXMSMSMMXMXAXAMXSMMSXAMMMAAAXASASMSSSMMXSXMAXXXXMMXMAMXAMAAMXMAAASMMMAMSMSAMXXAMXSASMSMMSSMMSASMMASMMXMSMXAMMMSXMXAMXMASXAMXSAMXSASA\nAMMMMAMAMMXMAXAMSXMSMSMMMMAXXMSSSMMSSXMXMASAMSAMXMMSSSSMMSAMAMMAMXMSAMAMAXAMMXMXXXMMSSMSXSASAASAMXMAAAMAAXMXSAMSAMXMAXMMSMMMAXAMMMAMSMMXMASM\nXMAAMASAMMAMAMAMAMXMAMXAMMAMAXXAMSAMMMMXMXMAMSMSAMXMAXAAASXSAXSAMAXSASASASXMXSSMMMSAAAAXXMMSXXSMMMMSSSMXSSMAMASXAMXSSSXAAAASMMMSAXSAMXMASAMX\nMMSSSMSMSMXMASAMMXAMAMMXSXXAASMAMMXMSAMSSXSAMXAAMXXSSSMMMSAMAXMASXMXMSASAAAMMSAAMAMMSMXMXMAXMMXXSXAAAAXXXAMAMXMXXMXMMAMSSMMMXSAAMMXXMMSMSXMX\nAXAAMXSAAMXMASXSXMXSAMXXMMMXMXMAMMMMMSAAXMASAMXMASXAXAXMXMXMSMSXMMMXXXAMXMASXSMMMSAMXMASMMXSAASAMMSSSMMMSXMXXMAMXMASXAXAMXSXXMAXMMMSMMAMMASX\nSAMAMAMSMMAMXSXMASAMXSXMSASMMXSSSMAXAMMAMXSAMXXXXSAMXMXAAMAXMAMAAASMSXSAMSAMMSMSXMASASMSAXMAMSMAXAXMMXAMMASAXMAMXSASMMXMSASAXSMSMSMAAMXXSAMA\nAMXXMMMAXSSSXMMMAMMSMAAASMMAMAMMAMSSXXXAMXMASMAMMMAXAXMSASAXSAMMSMSAMXXAMMMSAXAXXMASAMAXMSXMMXMMMSSMXXMAMSMASMMMMMASASAXMASAXXAAAMSSSMSMMASM\nMXMASMSMMXMAXAAXXSXAMMXMSASMMAXXAMXAMMMMMXSAMMAMXSASXSAXMAMXSAMXMXMAMMMAMAAMXMSMASAMMMSMXMAXSXXMAMAXMASMSAMAMXAAASXMAMXXMXMXSAMMSMAXMAAAAAMX\nAAAMMAAXMSMMMSSMMMXSSMXXXMMXSASMSSSMAAAAXMMASXXSAXXSMXMXMSMMXSMAXASAMAXXMMMSAMMAXMASMMMAAXSMSAMMSSMMAMAXAAMASMSSXSAMSMMMSMSAXAMAXMSMMSMMMSSX\nSXXXMSMXMMASMXMMMAAMAAMSSMMAMMMAMAAMXSXXMAMAMMAMXMAXAXMSMMAMSMMXSMSXSMMAMXXAXMXXXSMMMAMXSXXAXMXMXAXXSXXXSXMXXMMMXMXMAAXSAAMXXAMXMAXAXAAXSMMM\nAMMSMMXSASAMMASMMMSXMMMAAAMASAMSMSMMAXXMMSMMAMSMAAMXSXMAXMAMAAMXMMMMAMMMMMSMMMMMMMSASMSAMMMSMXMMSMMMMMMMXMASMSSMSSMSSSMSMSSMMSSSSSSSMMSMXAAS\nMAXMAAASXMXSMAMXAAMXSAMSSMMXMAXXAXXMASASAAAMSAMXXSXAAAAMXSASMSMAAAXMASAAAAAAAXAXAAMMSAMASAMASASAAAAAAAAMASAMXAAAAMMAMAAXMAAAMAAXAXAAXAAMSMMS\nSSXSMMMSXXASMMSSMXSASAXMMMXSAMXMSMMMSSMMSSSMXMASAAMXSMMXMAXMXXMSSMXSASMSMSSSMSXSMXMMMAMSMXSXSSMSSSXSSSSSMSAXMSMMMSMASMMMMSSMMMSMXMAMMMSXXAMX"
	var grid [][]string
	for _, row := range strings.Split(input, "\n") {
		grid = append(grid, strings.Split(row, ""))
	}
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
