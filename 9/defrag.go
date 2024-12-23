package main

import (
	"advent24/common"
	"slices"
	"strconv"
	"strings"
)

func main() {
	input := common.GetInput("9")
	var unfoldedInputArray []string
	nextFileId := 0
	for i := 0; i < len(input); i += 1 {
		character := ""
		if i%2 == 0 {
			character = strconv.Itoa(nextFileId)
			nextFileId++
		} else {
			character = "."
		}
		blockLength, _ := strconv.Atoi(string(input[i]))
		for j := 0; j < blockLength; j++ {
			unfoldedInputArray = append(unfoldedInputArray, character)
		}
	}
	for isFragmented(unfoldedInputArray) {
		earliestFreeIndex := slices.Index(unfoldedInputArray, ".")
		for i := len(unfoldedInputArray) - 1; i >= 0; i-- {
			if unfoldedInputArray[i] != "." {
				unfoldedInputArray[earliestFreeIndex] = unfoldedInputArray[i]
				unfoldedInputArray[i] = "."
				break
			}
		}
	}
	checksum := 0
	println(strings.Join(unfoldedInputArray, ""))
	for i, char := range unfoldedInputArray {
		if char == "" {
			break
		}
		col, _ := strconv.Atoi(char)
		checksum += i * col
	}
	println(checksum)

	var fsMap [][]string
	nextFileId = 0
	for i := 0; i < len(input); i += 1 {
		character := ""
		if i%2 == 0 {
			character = strconv.Itoa(nextFileId)
			nextFileId++
		} else {
			character = "."
		}
		blockLength, _ := strconv.Atoi(string(input[i]))
		blockList := []string{}
		for j := 0; j < blockLength; j++ {
			blockList = append(blockList, character)
		}
		fsMap = append(fsMap, blockList)
	}
	isFragmented := true
	for isFragmented {
		itemMoved := false
		for i := len(fsMap) - 1; i >= 0; i-- {
			if len(fsMap[i]) == 0 || fsMap[i][0] == "." {
				continue
			}
			if i == 0 {
				isFragmented = false
			}
			blockLength := len(fsMap[i])
			for j := 0; j < i; j++ {
				if len(fsMap[j]) >= blockLength && fsMap[j][0] == "." {
					itemMoved = true
					movingItems := make([]string, blockLength)
					copy(movingItems, fsMap[i])
					fsMap[j] = fsMap[j][:len(fsMap[j])-blockLength]
					for i2 := range fsMap[i] {
						fsMap[i][i2] = "."
					}
					fsMap = append(fsMap[:j], append([][]string{movingItems}, fsMap[j:]...)...)
					break
				}
			}
			if itemMoved {
				break
			}
		}
	}
	var flattenedFsMap []string
	for _, block := range fsMap {
		flattenedFsMap = append(flattenedFsMap, block...)
	}
	println(strings.Join(flattenedFsMap, ""))
	checksum = 0
	for i, char := range flattenedFsMap {
		if char == "" {
			break
		}
		col, _ := strconv.Atoi(char)
		checksum += i * col
	}
	println(checksum)

}

func isFragmented(input []string) bool {
	parsedInput := strings.Split(strings.Join(input, ""), ".")
	if len(parsedInput) == 1 {
		return false
	}
	for i := 1; i < len(parsedInput); i++ {
		if parsedInput[i] != "" {
			return true
		}
	}
	return false
}
