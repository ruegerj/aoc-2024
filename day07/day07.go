package day07

import (
	"math"
	"strings"

	"github.com/ruegerj/aoc-2024/pkg/common"
	"github.com/ruegerj/aoc-2024/pkg/util"
)

// Credits for optimization: https://www.reddit.com/r/adventofcode/comments/1h8l3z5/comment/m0tv6di/

type Day07 struct{}

func (d Day07) Part1(input string) *common.Solution {
	calibrations := parseCalibrations(input)
	totalCalibrationResult := 0

	for _, calibration := range calibrations {
		if isTractable(calibration.target, calibration.equationNumbers, false) {
			totalCalibrationResult += calibration.target
		}
	}

	return common.NewSolution(1, totalCalibrationResult)
}

func (d Day07) Part2(input string) *common.Solution {
	calibrations := parseCalibrations(input)
	totalCalibrationResult := 0

	for _, calibration := range calibrations {
		if isTractable(calibration.target, calibration.equationNumbers, true) {
			totalCalibrationResult += calibration.target
		}
	}

	return common.NewSolution(2, totalCalibrationResult)
}

type calibration struct {
	target          int
	equationNumbers []int
}

func endsWith(a, b int) bool {
	return math.Pow(float64((a-b)%10), float64(util.Digits(b))) == 0
}

func removeLast(numbers []int) ([]int, int) {
	if len(numbers) <= 0 {
		return []int{}, -1
	}

	lastIndex := len(numbers) - 1
	head := numbers[:lastIndex]
	last := numbers[lastIndex]

	return head, last
}

func isTractable(target int, numbers []int, checkConcat bool) bool {
	head, last := removeLast(numbers)
	if len(head) == 0 {
		return last == target
	}

	// check if division could be derived from remaining
	quotient, remainder := util.DivMod(target, last)
	if remainder == 0 && isTractable(quotient, head, checkConcat) {
		return true
	}

	// check if concatenation could be derived from remaining
	expectedTarget := target / int(math.Pow(10, float64(util.Digits(last))))
	if checkConcat && endsWith(target, last) && isTractable(expectedTarget, head, checkConcat) {
		return true
	}

	// default: check if plus operation could be derived from remaining
	return isTractable(target-last, head, checkConcat)

}

func parseCalibrations(input string) []calibration {
	lines := util.Lines(input)
	calibrations := make([]calibration, len(lines))

	for i, line := range lines {
		parts := strings.Split(line, ": ")
		calibration := calibration{
			target:          util.MustParseInt(parts[0]),
			equationNumbers: make([]int, 0),
		}
		for _, v := range strings.Split(parts[1], " ") {
			calibration.equationNumbers = append(calibration.equationNumbers, util.MustParseInt(v))
		}

		calibrations[i] = calibration
	}
	return calibrations
}
