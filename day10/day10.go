package day10

import (
	"maps"
	"strings"

	"github.com/ruegerj/aoc-2024/pkg/common"
	"github.com/ruegerj/aoc-2024/pkg/util"
)

type Day10 struct{}

func (d Day10) Part1(input string) *common.Solution {
	grid, trailHeads := parseMap(input)
	totalScore := 0

	for _, th := range trailHeads {
		found := make(map[point]int)
		wanderTrail(th.x, th.y, grid, found)
		totalScore += len(found)
	}

	return common.NewSolution(1, totalScore)
}

func (d Day10) Part2(input string) *common.Solution {
	grid, trailHeads := parseMap(input)
	totalScore := 0

	for _, th := range trailHeads {
		found := make(map[point]int)
		wanderTrail(th.x, th.y, grid, found)
		score := 0
		for visitCount := range maps.Values(found) {
			score += visitCount
		}
		totalScore += score
	}

	return common.NewSolution(2, totalScore)
}

type point struct {
	x, y int
}

func wanderTrail(cx, cy int, grid [][]int, found map[point]int) {
	cur := grid[cy][cx]
	if cur == 9 {
		found[point{x: cx, y: cy}]++
		return
	}

	if cx+1 < len(grid[0]) && grid[cy][cx+1]-cur == 1 {
		wanderTrail(cx+1, cy, grid, found)
	}
	if cx-1 >= 0 && grid[cy][cx-1]-cur == 1 {
		wanderTrail(cx-1, cy, grid, found)
	}
	if cy+1 < len(grid) && grid[cy+1][cx]-cur == 1 {
		wanderTrail(cx, cy+1, grid, found)
	}
	if cy-1 >= 0 && grid[cy-1][cx]-cur == 1 {
		wanderTrail(cx, cy-1, grid, found)
	}
}

func parseMap(input string) ([][]int, []point) {
	lines := util.Lines(input)
	trailHeads := make([]point, 0)
	grid := make([][]int, len(lines))
	for row, l := range lines {
		grid[row] = make([]int, len(l))
		for col, c := range strings.Split(l, "") {
			grid[row][col] = util.MustParseInt(c)
			if c == "0" {
				trailHeads = append(trailHeads, point{x: col, y: row})
			}
		}
	}
	return grid, trailHeads
}
