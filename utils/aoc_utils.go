package aoc_utils

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

type Problem struct {
	Day   int
	Part1 Part
	Part2 Part
}

type Part struct {
	InputFileLoc string
	InputString  string
	Result       int
}

func (p Part) OpenInputFile() *os.File {
	file, err := os.Open(p.InputFileLoc)
	if err != nil {
		fmt.Errorf("failed to load input file")
	}

	return file
}

func (p *Part) ReadInputFile() {
	file, err := os.Open(p.InputFileLoc)
	if err != nil {
		fmt.Errorf("Failed to load input")
	}

	scanner := bufio.NewScanner(file)
	input := ""
	for scanner.Scan() {
		input += scanner.Text()
	}
	p.InputString = input
	// return input
}

func (p Problem) PrintResults() {
	fmt.Printf("-- Day %d --\n", p.Day)
	fmt.Printf("Part 1: %v\nPart 2: %v\n", p.Part1.Result, p.Part2.Result)
}

// TODO: Need to figure out how to better handle problems that require mods between part 1 and part 2
func NewProblem(day int) Problem {
	problem := Problem{
		Day: day,
		Part1: Part{
			InputFileLoc: fmt.Sprintf("inputs/day%d.txt", day),
		},
		Part2: Part{
			InputFileLoc: fmt.Sprintf("inputs/day%d-part2.txt", day),
		},
	}
	problem.Part1.ReadInputFile()
	problem.Part2.ReadInputFile()

	return problem
}

func AbsInt(x int) int {
	if x < 0 {
		x = -x
	}

	return x
}

func IsOrdered(arr []int) bool {
	clone := slices.Clone(arr)
	if slices.IsSorted(clone) == false {
		slices.Reverse(clone)
	}

	return slices.IsSorted(clone)
}
