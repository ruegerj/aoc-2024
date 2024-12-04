package day04

import (
	"slices"
	"strings"

	"github.com/ruegerj/aoc-2024/pkg/common"
	"github.com/ruegerj/aoc-2024/pkg/util"
)

type Day04 struct{}

type direction int

const (
	UP = iota
	DOWN
	LEFT
	RIGHT
	UP_LEFT
	UP_RIGHT
	DOWN_LEFT
	DOWN_RIGHT
)

func (d Day04) Part1(input string) *common.Solution {
	const xmas = "XMAS"
	matrix := util.Matrix(input, "")

	directions := []direction{UP, DOWN, LEFT, RIGHT, UP_LEFT, UP_RIGHT, DOWN_LEFT, DOWN_RIGHT}
	width := len(matrix[0])
	height := len(matrix)
	xmasCount := 0

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			char := matrix[y][x]
			if char != string(xmas[0]) {
				continue
			}

			for _, d := range directions {
				if matchWord(x, y, d, xmas, matrix) {
					xmasCount++
				}
			}
		}
	}

	return common.NewSolution(1, xmasCount)
}

func (d Day04) Part2(input string) *common.Solution {
	const mas = "MAS"

	matrix := util.Matrix(input, "")
	height := len(matrix)
	width := len(matrix[0])
	masCount := 0

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			char := matrix[y][x]
			if char != "A" {
				continue
			}

			if y <= 0 || y >= height-1 {
				continue
			}

			if x <= 0 || x >= width-1 {
				continue
			}

			tl := matrix[y-1][x-1]
			tr := matrix[y-1][x+1]
			bl := matrix[y+1][x-1]
			br := matrix[y+1][x+1]
			mid := matrix[y][x]

			if isBiDirectionalWord(mas, []string{tl, mid, br}) && isBiDirectionalWord(mas, []string{tr, mid, bl}) {
				masCount++
			}
		}
	}

	return common.NewSolution(2, masCount)
}

func (d direction) getVector() (int, int) {
	switch d {
	case UP:
		return 0, -1
	case DOWN:
		return 0, 1
	case LEFT:
		return -1, 0
	case RIGHT:
		return 1, 0
	case UP_LEFT:
		return -1, -1
	case UP_RIGHT:
		return 1, -1
	case DOWN_LEFT:
		return -1, 1
	case DOWN_RIGHT:
		return 1, 1
	}
	return 0, 0
}

func isBiDirectionalWord(word string, chars []string) bool {
	if strings.Join(chars, "") == word {
		return true
	}
	slices.Reverse(chars)
	return strings.Join(chars, "") == word
}

func matchWord(startX, startY int, dir direction, word string, matrix [][]string) bool {
	col := startX
	row := startY
	dx, dy := dir.getVector()
	chars := make([]string, 4)
	for i := 1; i < len(word); i++ {
		col += dx
		row += dy

		if col < 0 || col >= len(matrix[0]) {
			return false
		}

		if row < 0 || row >= len(matrix) {
			return false
		}

		if matrix[row][col] != string(word[i]) {
			return false
		}
		chars[i] = matrix[row][col]
	}

	return true
}
