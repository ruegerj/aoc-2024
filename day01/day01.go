package day01

import (
	"slices"
	"strings"

	"github.com/ruegerj/aoc-2024/pkg/common"
	"github.com/ruegerj/aoc-2024/pkg/util"
)

type Day01 struct{}

func (d Day01) Part1(input string) *common.Solution {
	lines := util.Lines(input)
	right := []int{}
	left := []int{}

	for _, line := range lines {
		pair := strings.Split(line, "   ")
		left = append(left, util.MustParseInt(pair[0]))
		right = append(right, util.MustParseInt(pair[1]))
	}

	slices.Sort(left)
	slices.Sort(right)

	diff := 0
	for i := 0; i < len(left); i++ {
		diff += util.Abs(left[i] - right[i])
	}

	return common.NewSolution(1, diff)
}

func (d Day01) Part2(input string) *common.Solution {
	lines := util.Lines(input)
	left := []int{}
	right := map[int]int{}

	for _, line := range lines {
		pair := strings.Split(line, "   ")
		left = append(left, util.MustParseInt(pair[0]))
		r := util.MustParseInt(pair[1])
		right[r] = right[r] + 1
	}

	score := 0
	for _, l := range left {
		rightCount := right[l]
		score += l * rightCount
	}
	return common.NewSolution(2, score)
}
