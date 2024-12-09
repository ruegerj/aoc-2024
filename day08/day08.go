package day08

import (
	"github.com/ruegerj/aoc-2024/pkg/common"
	"github.com/ruegerj/aoc-2024/pkg/util"
)

type Day08 struct{}

func (d Day08) Part1(input string) *common.Solution {
	antennas, grid := parseAntennas(input)
	antiNodes := make(map[antenna]bool)

	for _, ant := range antennas {
		for _, other := range antennas {
			if ant == other || !ant.isSameType(other) {
				continue
			}

			dx := ant.x - other.x
			dy := ant.y - other.y
			dxi := dx * -1
			dyi := dy * -1

			antinode1 := newAntiNode(ant.x+dx, ant.y+dy)
			antinode2 := newAntiNode(other.x+dxi, other.y+dyi)

			if antinode1.isInBounds(grid) {
				antiNodes[antinode1] = true
			}
			if antinode2.isInBounds(grid) {
				antiNodes[antinode2] = true
			}
		}
	}

	uniqueAntiNodesCount := len(antiNodes)

	return common.NewSolution(1, uniqueAntiNodesCount)
}

func (d Day08) Part2(input string) *common.Solution {
	antennas, grid := parseAntennas(input)
	antiNodes := make(map[antenna]bool)

	for _, ant := range antennas {
		for _, other := range antennas {
			if ant == other || !ant.isSameType(other) {
				continue
			}

			dx := ant.x - other.x
			dy := ant.y - other.y
			dxi := dx * -1
			dyi := dy * -1

			nextAntiNode := newAntiNode(ant.x+dx, ant.y+dy)
			for nextAntiNode.isInBounds(grid) {
				antiNodes[nextAntiNode] = true
				nextAntiNode = newAntiNode(nextAntiNode.x+dx, nextAntiNode.y+dy)
			}

			nextAntiNode = newAntiNode(ant.x+dxi, ant.y+dyi)
			for nextAntiNode.isInBounds(grid) {
				antiNodes[nextAntiNode] = true
				nextAntiNode = newAntiNode(nextAntiNode.x+dxi, nextAntiNode.y+dyi)
			}
		}
	}

	uniqueAntiNodesCount := len(antiNodes)

	return common.NewSolution(2, uniqueAntiNodesCount)
}

type antenna struct {
	x, y   int
	symbol string
}

func newAntenna(symbol string, x, y int) antenna {
	return antenna{symbol: symbol, x: x, y: y}
}

func newAntiNode(x, y int) antenna {
	return antenna{symbol: "#", x: x, y: y}
}

func (a antenna) isSameType(other antenna) bool {
	return a.symbol == other.symbol
}

func (a antenna) isInBounds(grid [][]string) bool {
	return a.x >= 0 && a.x < len(grid[0]) && a.y >= 0 && a.y < len(grid)
}

func parseAntennas(input string) ([]antenna, [][]string) {
	grid := util.Matrix(input, "")
	antennas := make([]antenna, 0)
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[0]); x++ {
			char := grid[y][x]
			if char == "." {
				continue
			}
			antennas = append(antennas, newAntenna(char, x, y))
		}
	}
	return antennas, grid
}
