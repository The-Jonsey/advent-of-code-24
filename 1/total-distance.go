package main

import (
	"advent24/common"
	"math"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

var splitOnStringRegex, _ = regexp.Compile("\\s+")

func convertToSortedIntLists(items []string) ([]int, []int) {
	left := make([]int, len(items))
	right := make([]int, len(items))
	for i, row := range items {
		rowSplit := splitOnStringRegex.Split(row, -1)
		leftValue, _ := strconv.Atoi(rowSplit[0])
		rightValue, _ := strconv.Atoi(rowSplit[1])
		left[i] = leftValue
		right[i] = rightValue
	}
	sort.Ints(left)
	sort.Ints(right)
	return left, right
}

func calculateTotalDistance(left []int, right []int) int {
	totalDistance := 0
	for index, leftValue := range left {
		rightValue := right[index]
		totalDistance += int(math.Abs(float64(leftValue - rightValue)))
	}
	return totalDistance
}

func calculateSimilarityScore(left []int, right []int) int {
	totalScore := 0
	for _, leftValue := range left {
		localTotal := 0
		for _, rightValue := range right {
			if rightValue > leftValue {
				break
			}
			if rightValue == leftValue {
				localTotal += 1
			}
		}
		totalScore += leftValue * localTotal
	}
	return totalScore
}

func main() {
	input := common.GetInput("1")
	left, right := convertToSortedIntLists(strings.Split(input, "\n"))
	println(calculateTotalDistance(left, right))
	println(calculateSimilarityScore(left, right))
}
