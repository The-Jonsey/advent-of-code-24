package main

import (
	"advent24/common"
	"log"
	"regexp"
	"strconv"
)

var validMulRegex, _ = regexp.Compile(`mul\(\d{1,3},\d{1,3}\)`)
var validMulDontDoRegex, _ = regexp.Compile(`(mul\(\d{1,3},\d{1,3}\)|don't\(\)|do\(\))`)
var digitRegex, _ = regexp.Compile(`\d{1,3}`)

func calculateMul(match string) int {
	digits := digitRegex.FindAllString(match, -1)
	first, err := strconv.Atoi(digits[0])
	if err != nil {
		log.Fatalln(err)
	}
	second, err := strconv.Atoi(digits[1])
	if err != nil {
		log.Fatalln(err)
	}
	return first * second
}

func main() {
	input := common.GetInput("3")
	matches := validMulRegex.FindAllString(input, -1)
	output := 0
	for _, match := range matches {
		output += calculateMul(match)
	}
	println(output)
	matches = validMulDontDoRegex.FindAllString(input, -1)
	output = 0
	isEnabled := true
	for _, match := range matches {
		if match == "don't()" {
			isEnabled = false
			continue
		} else if match == "do()" {
			isEnabled = true
			continue
		}
		if isEnabled {
			output += calculateMul(match)
		}
	}
	println(output)
}
