package main

import (
	"advent24/common"
	"math"
	"strconv"
	"strings"
)

func indexOf(item string, list []string) int {
	for i := range list {
		if list[i] == item {
			return i
		}
	}
	return -1
}

func passesRule(first string, second string, list []string) bool {
	indexOfFirst := indexOf(first, list)
	if indexOfFirst == -1 {
		return true
	}
	indexOfSecond := indexOf(second, list)
	if indexOfSecond == -1 {
		return true
	}
	return indexOfSecond > indexOfFirst
}

func main() {
	rulesInput := common.GetInput("5/rules")
	rulesSplit := strings.Split(rulesInput, "\n")
	rules := make([][]string, len(rulesSplit)-1)
	for i, rule := range rulesSplit {
		ruleSplit := strings.Split(rule, "|")
		if len(ruleSplit) == 2 {
			rules[i] = ruleSplit
		}
	}
	updatesInput := common.GetInput("5/updates")
	middlePageNumberSum := 0
	updates := strings.Split(updatesInput, "\n")
	var malformedUpdates [][]string
	for _, update := range updates {
		if update == "" {
			continue
		}
		updateList := strings.Split(update, ",")
		passesAllRules := true
		for _, rule := range rules {
			if !passesRule(rule[0], rule[1], updateList) {
				passesAllRules = false
			}
		}
		if !passesAllRules {
			malformedUpdates = append(malformedUpdates, updateList)
			continue
		}
		middlePageIndex := int(math.Floor(float64(len(updateList)) / 2.0))
		middlePageNumber, _ := strconv.Atoi(updateList[middlePageIndex])
		middlePageNumberSum += middlePageNumber
	}
	println(middlePageNumberSum)
	middlePageNumberSum = 0

	for i := 0; i < len(malformedUpdates); i++ {
		update := malformedUpdates[i]
		requiredUpdate := false
		for _, rule := range rules {
			if !passesRule(rule[0], rule[1], update) {
				update = moveElement(update, indexOf(rule[0], update), indexOf(rule[1], update))
				requiredUpdate = true
			}
		}
		if requiredUpdate {
			i -= 1
			continue
		}
		middlePageIndex := int(math.Floor(float64(len(update)) / 2.0))
		middlePageNumber, _ := strconv.Atoi(update[middlePageIndex])
		middlePageNumberSum += middlePageNumber
	}
	println(middlePageNumberSum)
}

func moveElement(slice []string, sourceIndex int, destinationIndex int) []string {
	if sourceIndex < 0 || sourceIndex >= len(slice) || destinationIndex < 0 || destinationIndex > len(slice) {
		panic("indices out of range")
	}

	// Extract the element to move
	element := slice[sourceIndex]

	// Remove the element from the slice
	slice = append(slice[:sourceIndex], slice[sourceIndex+1:]...)

	// Adjust destinationIndex if it is after the sourceIndex
	if destinationIndex > sourceIndex {
		destinationIndex--
	}

	// Insert the element at the new position
	slice = append(slice[:destinationIndex], append([]string{element}, slice[destinationIndex:]...)...)

	return slice
}
