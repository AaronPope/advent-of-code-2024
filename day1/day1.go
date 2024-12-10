package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"

	utils "aoc2024/utils"
)

func TestPub() {
	fmt.Println("TESTING...")
}

func main() {
	file, err := os.Open("inputs/day1.txt")

	if err != nil {
		fmt.Errorf("Failed to load input")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	left := make([]int, 1)
	right := make([]int, 1)
	for scanner.Scan() {
		line := string(scanner.Bytes())
		nums := strings.Split(line, "   ")

		l, _ := strconv.Atoi(nums[0])
		r, _ := strconv.Atoi(nums[1])
		left = append(left, l)
		right = append(right, r)
	}

	slices.Sort(left)
	slices.Sort(right)

	sum := 0

	for i := 0; i < len(left); i++ {
		diff := left[i] - right[i]
		diff = utils.AbsInt(diff)
		sum += diff
	}

	fmt.Println(sum) // Confirmed correct: 2375403
}
