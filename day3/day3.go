package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var digitsRE = regexp.MustCompile(`\d+`)
var re = regexp.MustCompile(`mul\(\d+,\d+\)`)

// TODO: Refactor + add a test for funsies
func main() {
	// sample := "1532mul(5,4)k,mu (l,5)testingmul(2,2)testing"
	// res := re.FindAll([]byte(sample), -1)
	// for _, v := range res {
	// 	fmt.Println(string(v))
	// }

	file, err := os.Open("inputs/day3.txt")

	if err != nil {
		fmt.Errorf("Failed to load input")
	}
	defer file.Close()

	part1 := Part1(file)
	part2 := Part2()
	fmt.Printf("Part 1: %v\nPart 2: %v\n", part1, part2)
}

func Part1(input *os.File) int {
	scanner := bufio.NewScanner(input)
	product := 0
	for scanner.Scan() {
		chunk := scanner.Text()

		muls := re.FindAllString(chunk, -1)

		for _, v := range muls {
			nums := digitsRE.FindAllString(v, -1)

			n1, _ := strconv.Atoi(nums[0])
			n2, _ := strconv.Atoi(nums[1])

			product += n1 * n2
		}
	}
	return product
}

func Part2() int {
	return 0
}
