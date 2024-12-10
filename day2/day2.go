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

	safeReports := 0

	for scanner.Scan() {
		line := string(scanner.Bytes())
		nums := strings.Split(line, " ")

		ints := make([]int, 0)
		for _, n := range nums {
			i, _ := strconv.Atoi(n)
			ints = append(ints, i)
		}

		if utils.IsOrdered(ints) && checkElementGaps(ints) {
			safeReports++
		}
	}

	fmt.Printf("Safe Reports: %d", safeReports)
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
