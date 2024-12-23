package main

import (
	"advent24/common"
	"strconv"
	"strings"
)

func main() {
	input := common.GetInput("7")
	inputLines := strings.Split(input, "\n")
	actualValue := 0
	for _, line := range inputLines {
		lineSplit := strings.Split(line, ": ")
		expectedValue := lineSplit[0]
		intExpectedValue, _ := strconv.Atoi(expectedValue)
		inputValues := strings.Split(lineSplit[1], " ")
		var intInputValues []int
		for _, value := range inputValues {
			intValue, _ := strconv.Atoi(value)
			intInputValues = append(intInputValues, intValue)
		}
		if len(intInputValues) == 1 && intInputValues[0] == intExpectedValue {
			actualValue += intExpectedValue
			continue
		}
		if isValidInputs(intInputValues, intExpectedValue, 1, intInputValues[0]) {
			actualValue += intExpectedValue
		}
	}
	println(actualValue)
}

func isValidInputs(inputs []int, target int, index int, currentValue int) bool {
	if index == len(inputs) {
		return currentValue == target
	}
	nextValue := inputs[index]
	concatenatedStr := strconv.Itoa(currentValue) + strconv.Itoa(nextValue)
	concatInts, _ := strconv.Atoi(concatenatedStr)
	return isValidInputs(inputs, target, index+1, currentValue+nextValue) ||
		isValidInputs(inputs, target, index+1, currentValue*nextValue) ||
		isValidInputs(inputs, target, index+1, concatInts)
}
