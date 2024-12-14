package day1

import (
	"bufio"
	"slices"
	"strconv"
	"strings"

	utils "aoc2024/utils"
)

func main() {
	Solve()
}

func Solve() utils.Problem {
	p := utils.NewProblem(1)
	file := p.Part1.OpenInputFile()
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

	p.Part1.Result = sortedDiffs(left, right)
	p.Part2.Result = occurrencesProduct(left, right)

	return p
}

func sortedDiffs(left []int, right []int) int {
	slices.Sort(left)
	slices.Sort(right)

	sum := 0

	for i := 0; i < len(left); i++ {
		diff := left[i] - right[i]
		diff = utils.AbsInt(diff)
		sum += diff
	}

	return sum
}

func occurrencesProduct(left []int, right []int) int {
	rightOccurrences := make(map[int]int)

	for _, n := range right {
		_, exists := rightOccurrences[n]

		if exists == false {
			rightOccurrences[n] = 1
		} else {
			rightOccurrences[n] = rightOccurrences[n] + 1
		}
	}

	total := 0
	for _, n := range left {
		total += n * rightOccurrences[n]
	}

	return total
}
