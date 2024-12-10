package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	utils "aoc2024/utils"
)

func main() {
	file, err := os.Open("inputs/day2.txt")

	if err != nil {
		fmt.Errorf("Failed to load input")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	input := make([][]int, 0)
	for scanner.Scan() {
		line := string(scanner.Bytes())
		nums := strings.Split(line, " ")

		ints := make([]int, 0)
		for _, n := range nums {
			i, _ := strconv.Atoi(n)
			ints = append(ints, i)
		}

		input = append(input, ints)
	}

	part1 := Part1(input)
	part2 := Part2(input)

	fmt.Printf("Part 1: %d\n", part1)
	fmt.Printf("Part 2: %d", part2)
}

func Part1(input [][]int) int {
	safeReports := 0

	for _, arr := range input {
		if utils.IsOrdered(arr) && checkElementGaps(arr) {
			safeReports++
		}
	}

	return safeReports
}

func Part2(input [][]int) int {
	safeReports := 0
	for _, arr := range input {
		if utils.IsOrdered(arr) && checkElementGaps(arr) {
			safeReports++
		} else {
			for i, _ := range arr {
				left := arr[0:i]
				right := arr[i+1:]

				appended := make([]int, 0)
				appended = append(appended, left...)
				appended = append(appended, right...)

				if utils.IsOrdered(appended) && checkElementGaps(appended) {
					safeReports++
					break
				}
			}
		}
	}

	// TODO
	return safeReports
}

func checkElementGaps(arr []int) bool {
	for i := 0; i < len(arr)-1; i++ {
		diff := utils.AbsInt(arr[i] - arr[i+1])

		if diff == 0 || diff > 3 {
			return false
		}
	}

	return true
}
