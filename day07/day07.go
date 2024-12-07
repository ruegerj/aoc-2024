package day07

import (
	"math"
	"slices"
	"strings"

	"github.com/ruegerj/aoc-2024/pkg/common"
	"github.com/ruegerj/aoc-2024/pkg/util"
)

type Day07 struct{}

func (d Day07) Part1(input string) *common.Solution {
	calibrations := parseCalibrations(input)
	var totalCalibrationResult int64 = 0

	for _, calibration := range calibrations {
		if calibration.tryToSolve([]rune{'+', '*'}) {
			totalCalibrationResult += calibration.target
		}
	}

	return common.NewSolution(1, totalCalibrationResult)
}

func (d Day07) Part2(input string) *common.Solution {
	calibrations := parseCalibrations(input)
	var totalCalibrationResult int64 = 0

	for _, calibration := range calibrations {
		if calibration.tryToSolve([]rune{'+', '*', '|'}) {
			totalCalibrationResult += calibration.target
		}
	}

	return common.NewSolution(2, totalCalibrationResult)
}

type calibration struct {
	target          int64
	equationNumbers []int64
}

func (c calibration) tryToSolve(operators []rune) bool {
	usePlus := slices.Contains(operators, '+')
	useMultiply := slices.Contains(operators, '*')
	useConcat := slices.Contains(operators, '|')

	var opResult int64
	lastOp := int(math.Pow(3, float64(len(c.equationNumbers))) - 1)
	for i := 0; i < lastOp; i++ {
		opResult = c.equationNumbers[0]
		for j := 1; j < len(c.equationNumbers); j++ {
			currentOp := i
			for k := 1; k < j; k++ {
				currentOp /= 3
			}
			if currentOp%3 == 0 && usePlus {
				opResult += c.equationNumbers[j]
			} else if currentOp%3 == 1 && useMultiply {
				opResult *= c.equationNumbers[j]
			} else if useConcat {
				opResult = int64(util.Concat(int(opResult), int(c.equationNumbers[j])))
			}
		}
		if c.target == opResult {
			return true
		}
	}
	return false

}

func parseCalibrations(input string) []calibration {
	lines := util.Lines(input)
	calibrations := make([]calibration, len(lines))

	for i, line := range lines {
		parts := strings.Split(line, ": ")
		calibration := calibration{
			target:          util.MustParseInt64(parts[0]),
			equationNumbers: make([]int64, 0),
		}
		for _, v := range strings.Split(parts[1], " ") {
			calibration.equationNumbers = append(calibration.equationNumbers, util.MustParseInt64(v))
		}

		calibrations[i] = calibration
	}
	return calibrations
}
