package common

import (
	"log"
	"os"
	"strings"
)

func GetInput(subDir string) string {
	dat, err := os.ReadFile(subDir + "/input.txt")
	if err != nil {
		log.Fatalln(err)
	}
	return string(dat)
}

func InputTo2dArray(input string) [][]string {
	return InputTo2dArrayWithSep(input, "")
}

func InputTo2dArrayWithSep(input string, seperator string) [][]string {
	inputLines := strings.Split(input, "\n")
	var grid [][]string
	for _, line := range inputLines {
		grid = append(grid, strings.Split(line, seperator))
	}
	return grid
}
