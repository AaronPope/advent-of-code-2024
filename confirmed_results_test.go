package main

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"aoc2024/day1"
	"aoc2024/day2"
	"aoc2024/day3"
	"aoc2024/day4"
	utils "aoc2024/utils"
)

func CheckProblemResults(t *testing.T, p utils.Problem,
	expectedPart1 int, expectedPart2 int) {

	assert.Equal(t, expectedPart1, p.Part1.Result)
	assert.Equal(t, expectedPart2, p.Part2.Result)
}

func TestConfirmedResultsDay1(t *testing.T) {
	CheckProblemResults(t, day1.Solve(),
		2375403,
		23082277)
}

func TestConfirmedResultsDay2(t *testing.T) {
	CheckProblemResults(t, day2.Solve(),
		526,
		566)

}

func TestConfirmedResultsDay3(t *testing.T) {
	CheckProblemResults(t, day3.Solve(),
		190604937,
		82857512)
}

func TestConfirmedResultsDay4(t *testing.T) {
	CheckProblemResults(t, day4.Solve(),
		0,
		0)
}
