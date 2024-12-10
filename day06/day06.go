package day06

import (
	"fmt"
	"strings"

	"github.com/ruegerj/aoc-2024/pkg/common"
	"github.com/ruegerj/aoc-2024/pkg/util"
)

type Day06 struct{}

func (d Day06) Part1(input string) *common.Solution {
	current, grid := parseInput(input)
	path := make(map[*point]bool)
	var dir direction = up

	for true {
		dx, dy := dir.deltas()
		if !inBounds(current.x+dx, current.y+dy, grid) {
			path[current] = true
			break
		}

		next := grid[current.y+dy][current.x+dx]
		if next.value == "#" {
			dir = dir.rotateRight()
			continue
		}

		path[current] = true
		current = next
	}

	return common.NewSolution(1, len(path))
}

func (d Day06) Part2(input string) *common.Solution {
	current, grid := parseInput(input)
	origin := current
	loopStarts := make(map[string]bool)
	var dir direction = up

	for true {
		dx, dy := dir.deltas()
		if !inBounds(current.x+dx, current.y+dy, grid) {
			break
		}

		next := grid[current.y+dy][current.x+dx]
		if next.value == "#" {
			dir = dir.rotateRight()
			continue
		}

		next.value = "#"
		if current != origin && hasLoop(origin, up, grid) {
			loopStarts[pointKey(next)] = true
		}
		next.value = "."

		current = next
	}

	return common.NewSolution(2, len(loopStarts))
}

func hasLoop(cur *point, d direction, grid [][]*point) bool {
	path := make(map[string]bool)

	for true {
		// check if map has vector of point & direction
		if _, visited := path[vectorKey(cur, d)]; visited {
			return true
		}

		dx, dy := d.deltas()
		if !inBounds(cur.x+dx, cur.y+dy, grid) {
			return false
		}

		next := grid[cur.y+dy][cur.x+dx]
		if next.value == "#" {
			d = d.rotateRight()
			continue
		}

		path[vectorKey(cur, d)] = true
		cur = next
	}

	return false
}

func vectorKey(p *point, d direction) string {
	return fmt.Sprintf("%d,%d;%d", p.x, p.y, d)
}

func pointKey(p *point) string {
	return fmt.Sprintf("%d,%d", p.x, p.y)
}

type direction int

const (
	up = iota
	right
	down
	left
)

func (d direction) deltas() (int, int) {
	if d == up {
		return 0, -1
	}
	if d == right {
		return 1, 0
	}
	if d == down {
		return 0, 1
	}
	return -1, 0
}

func (d direction) rotateRight() direction {
	if d == up {
		return right
	}
	if d == right {
		return down
	}
	if d == down {
		return left
	}
	return up
}

type point struct {
	x, y  int
	value string
}

func inBounds(x, y int, grid [][]*point) bool {
	return x >= 0 && x < len(grid[0]) && y >= 0 && y < len(grid)
}

func parseInput(input string) (*point, [][]*point) {
	lines := util.Lines(input)
	var current *point
	grid := make([][]*point, len(lines))
	for row, l := range lines {
		grid[row] = make([]*point, len(l))
		for col, c := range strings.Split(l, "") {
			pnt := &point{x: col, y: row, value: c}
			if c == "^" {
				current = pnt
			}
			grid[row][col] = pnt
		}
	}

	return current, grid
}
