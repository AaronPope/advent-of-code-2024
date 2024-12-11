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

	part1 := Part1() // 190604937
	part2 := Part2() // 82857512
	fmt.Printf("Part 1: %v\nPart 2: %v\n", part1, part2)
}

func Part1() int {
	file, err := os.Open("inputs/day3.txt")

	if err != nil {
		fmt.Errorf("Failed to load input")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	input := ""
	for scanner.Scan() {
		input += scanner.Text()
	}

	product := 0

	muls := re.FindAllString(input, -1)

	for _, v := range muls {
		nums := digitsRE.FindAllString(v, -1)

		n1, _ := strconv.Atoi(nums[0])
		n2, _ := strconv.Atoi(nums[1])

		product += n1 * n2
	}

	return product
}

func Part2() int {
	// Had to add `do()` to the beginning and `don't()` to the end of the file
	//	since muls are active at the beginning and a `do()` is near the end but not closed by a `don't()`
	file, err := os.Open("inputs/day3-part2.txt")

	if err != nil {
		fmt.Errorf("Failed to load input")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	input := ""
	for scanner.Scan() {
		input += scanner.Text()
	}

	activeMulsRE := regexp.MustCompile(`do\(\)(.*?)don't\(\)`)
	activeSections := activeMulsRE.FindAllString(input, -1)

	product := 0
	for _, v := range activeSections {
		muls := re.FindAllString(v, -1)
		for _, v := range muls {
			nums := digitsRE.FindAllString(v, -1)

			n1, _ := strconv.Atoi(nums[0])
			n2, _ := strconv.Atoi(nums[1])

			product += n1 * n2
		}
	}

	return product
}
