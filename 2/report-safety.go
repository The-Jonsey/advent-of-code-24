package main

import (
	"advent24/common"
	"log"
	"math"
	"regexp"
	"strconv"
	"strings"
)

var splitOnStringRegex, _ = regexp.Compile("\\s+")

type LevelProgressionType int

const (
	INCREASING LevelProgressionType = iota
	DECREASING
	NONE
)

func getProgressionType(first int, second int) LevelProgressionType {
	if first < second {
		return INCREASING
	} else if first > second {
		return DECREASING
	}
	return NONE
}

func isSafe(levels []string, canBeDampened bool) bool {
	errorDetected := false
	firstValue, err := strconv.Atoi(levels[0])
	if err != nil {
		log.Fatalln("Error parsing int", firstValue)
	}
	secondValue, err := strconv.Atoi(levels[1])
	if err != nil {
		log.Fatalln("Error parsing int", secondValue)
	}
	progressionType := getProgressionType(firstValue, secondValue)
	if progressionType == NONE {
		errorDetected = true
	}
	for i := 0; i < len(levels)-1; i++ {
		firstValue, err = strconv.Atoi(levels[i])
		if err != nil {
			log.Fatalln("Error parsing int", firstValue)
		}
		secondValue, err = strconv.Atoi(levels[i+1])
		if err != nil {
			log.Fatalln("Error parsing int", secondValue)
		}
		if getProgressionType(firstValue, secondValue) != progressionType {
			errorDetected = true
		}
		if int(math.Abs(float64(firstValue-secondValue))) > 3 {
			errorDetected = true
		}
	}
	if errorDetected && canBeDampened {
		for i := range levels {
			dampenedLevels := append([]string{}, levels[0:i]...)
			dampenedLevels = append(dampenedLevels, levels[i+1:]...)
			if isSafe(dampenedLevels, false) {
				return true
			}
		}
		return false
	} else {
		return !errorDetected
	}
}

func main() {
	input := common.GetInput("2")
	rows := strings.Split(input, "\n")
	safeCount := 0
	for _, row := range rows {
		if isSafe(splitOnStringRegex.Split(row, -1), false) {
			safeCount += 1
		}
	}
	println(safeCount)
	safeCount = 0
	for _, row := range rows {
		if isSafe(splitOnStringRegex.Split(row, -1), true) {
			safeCount += 1
		}
	}
	println(safeCount)
}
