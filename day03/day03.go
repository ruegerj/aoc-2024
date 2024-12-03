package day03

import (
	"regexp"

	"github.com/ruegerj/aoc-2024/pkg/common"
	"github.com/ruegerj/aoc-2024/pkg/util"
)

type Day03 struct{}

func (d Day03) Part1(input string) *common.Solution {
	multiplications := parseMultiplications(input)
	sum := 0

	for _, multiplication := range multiplications {
		sum += multiplication.evaluate()
	}

	return common.NewSolution(1, sum)
}

func (d Day03) Part2(input string) *common.Solution {
	multiplications := parseMultiplications(input)
	instructions := parseInstructions(input)

	enabled := true
	hasNext := len(instructions) > 0
	nextInstrIdx := 0
	sum := 0

	for _, multiplication := range multiplications {
		if hasNext && multiplication.startPos > instructions[nextInstrIdx].startPos {
			enabled = instructions[nextInstrIdx].enabled()

			hasNext = len(instructions[nextInstrIdx:]) > 1
			if hasNext {
				nextInstrIdx++
			}
		}

		if enabled {
			sum += multiplication.evaluate()
		}
	}

	return common.NewSolution(2, sum)
}

type instruction struct {
	startPos int
	value    string
}

func (i instruction) enabled() bool {
	return i.value != "don't()"
}

type multiplication struct {
	startPos int
	left     int
	right    int
}

func (m multiplication) evaluate() int {
	return m.left * m.right
}

func parseMultiplications(input string) []multiplication {
	multiplyMatcher := regexp.MustCompile(`mul\((?P<left>\d+),(?P<right>\d+)\)`)
	multiplications := make([]multiplication, 0)
	multiplicationMatches := multiplyMatcher.FindAllStringSubmatchIndex(input, -1)
	for _, mm := range multiplicationMatches {
		multiplications = append(multiplications, multiplication{
			startPos: mm[0],
			left:     util.MustParseInt(input[mm[2]:mm[3]]),
			right:    util.MustParseInt(input[mm[4]:mm[5]]),
		})
	}
	return multiplications
}

func parseInstructions(input string) []instruction {
	instructionMatcher := regexp.MustCompile(`do(n't)?\(\)`)
	instructions := make([]instruction, 0)
	instructionMatches := instructionMatcher.FindAllStringIndex(input, -1)
	for _, im := range instructionMatches {
		instructions = append(instructions, instruction{
			startPos: im[0],
			value:    input[im[0]:im[1]],
		})
	}
	return instructions
}
