package day3

import (
	"regexp"
	"strconv"

	utils "aoc2024/utils"
)

var digitsRE = regexp.MustCompile(`\d+`)
var re = regexp.MustCompile(`mul\(\d+,\d+\)`)

func Solve() utils.Problem {
	p := utils.NewProblem(3)

	p.Part1.Result = Part1(p.Part1.InputString) // 190604937
	p.Part2.Result = Part2(p.Part2.InputString) // 82857512

	return p
}

func Part1(input string) int {
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

// TODO: Figure out a way to better handle this while leveraging the struct
func Part2(input string) int {
	// Had to add `do()` to the beginning and `don't()` to the end of the file
	//	since muls are active at the beginning and a `do()` is near the end but not closed by a `don't()`
	// file, err := os.Open("inputs/day3-part2.txt")

	// if err != nil {
	// 	fmt.Errorf("Failed to load input")
	// }
	// defer file.Close()
	// scanner := bufio.NewScanner(file)
	// input := ""
	// for scanner.Scan() {
	// 	input += scanner.Text()
	// }

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
