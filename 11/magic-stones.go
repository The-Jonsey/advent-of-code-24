package main

import (
	"advent24/common"
	"strconv"
	"strings"
	"sync"
)

func main() {
	input := common.GetInput("11")
	inputArr := strings.Split(input, " ")
	var stones []int
	for _, s := range inputArr {
		stone, _ := strconv.Atoi(s)
		stones = append(stones, stone)
	}
	//for i := 0; i < 75; i++ {
	//	println(i, ": ", len(stones))
	//	for j := 0; j < len(stones); j++ {
	//		stone := stones[j]
	//		if stone == 0 {
	//			stones[j] = 1
	//		} else if len(strconv.Itoa(stone))%2 == 0 {
	//			stoneString := strconv.Itoa(stone)
	//			newStone1, _ := strconv.Atoi(stoneString[:len(stoneString)/2])
	//			newStone2, _ := strconv.Atoi(stoneString[len(stoneString)/2 : len(stoneString)])
	//			newStonesList := []int{newStone1, newStone2}
	//			newStonesList = append(newStonesList, stones[j+1:len(stones)]...)
	//			stones = append(stones[:j], newStonesList...)
	//			j += 1
	//		} else {
	//			stones[j] = stones[j] * 2024
	//		}
	//	}
	//}
	//stoneCount := 0
	results := make(chan int, len(stones))
	var wg sync.WaitGroup
	for _, stone := range stones {
		wg.Add(1)
		go func(stone int) {
			defer wg.Done()
			results <- followStoneLifecycle(stone, 0, 75)
		}(stone)
	}
	wg.Wait()
	close(results)

	var stoneCount int
	for result := range results {
		stoneCount += result
	}
	println(stoneCount)
	//println(stoneCount)
}

func followStoneLifecycle(stone int, generation int, maxGenerations int) int {
	if generation == maxGenerations {
		return 1
	}
	if stone == 0 {
		return followStoneLifecycle(1, generation+1, maxGenerations)
	} else if len(strconv.Itoa(stone))%2 == 0 {
		stoneString := strconv.Itoa(stone)
		newStone1, _ := strconv.Atoi(stoneString[:len(stoneString)/2])
		newStone2, _ := strconv.Atoi(stoneString[len(stoneString)/2 : len(stoneString)])
		return followStoneLifecycle(newStone1, generation, maxGenerations) +
			followStoneLifecycle(newStone2, generation, maxGenerations)
	} else {
		return followStoneLifecycle(stone*2024, generation+1, maxGenerations)
	}
}
