package day02

import (
	"slices"
	"strings"

	"github.com/ruegerj/aoc-2024/pkg/common"
	"github.com/ruegerj/aoc-2024/pkg/util"
)

type Day02 struct{}

func (d Day02) Part1(input string) *common.Solution {
	reports := parseReports(input)
	safeCount := 0

	for _, report := range reports {
		if isSafeReport(report) {
			safeCount++
		}
	}

	return common.NewSolution(1, safeCount)
}

func (d Day02) Part2(input string) *common.Solution {
	reports := parseReports(input)
	safeCount := 0

	for _, report := range reports {
		safeWithError := false

		for i := 0; i < len(report); i++ {
			currentLevelRemoved := slices.Concat(report[0:i], report[i+1:])
			if isSafeReport(currentLevelRemoved) {
				safeWithError = true
				break
			}
		}

		if isSafeReport(report) || safeWithError {
			safeCount++
		}
	}

	return common.NewSolution(2, safeCount)
}

func isSafeReport(report []int) bool {
	diffs := []int{}

	for i := 1; i < len(report); i++ {
		diffs = append(diffs, report[i]-report[i-1])
	}

	increasing := util.Every(diffs, func(level int) bool {
		return level >= 1 && level <= 3
	})
	decreasing := util.Every(diffs, func(level int) bool {
		return level <= -1 && level >= -3
	})

	return increasing || decreasing
}

func parseReports(input string) [][]int {
	lines := util.Lines(input)
	reports := make([][]int, len(lines))
	for i, l := range lines {
		reports[i] = util.ToIntSlice(strings.Split(l, " "))
	}
	return reports
}
