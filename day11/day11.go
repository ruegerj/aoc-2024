package day11

import (
	"fmt"
	"maps"
	"strings"

	"github.com/ruegerj/aoc-2024/pkg/common"
	"github.com/ruegerj/aoc-2024/pkg/util"
)

type Day11 struct{}

const uneven_multiplier = 2024

func (d Day11) Part1(input string) *common.Solution {
	stones := parseStones(input)
	totalStones := blink(25, stones)
	return common.NewSolution(1, totalStones)
}

func (d Day11) Part2(input string) *common.Solution {
	stones := parseStones(input)
	totalStones := blink(75, stones)
	return common.NewSolution(2, totalStones)
}

func blink(times int, stones []int) int64 {
	totalStones := int64(len(stones))

	stoneCount := make(map[int64]int)
	for _, stone := range stones {
		stoneCount[int64(stone)]++
	}

	for i := 0; i < times; i++ {
		updatedCount := maps.Clone(stoneCount)

		for stone, count := range stoneCount {
			if count == 0 {
				continue
			}

			updatedCount[stone] -= count
			if stone == 0 {
				updatedCount[1] += count
				continue
			}

			digits := util.Digits64(stone)
			if digits%2 != 0 {
				newStone := stone * 2024
				updatedCount[newStone] += count
				continue
			}

			stringStone := fmt.Sprint(stone)
			left := util.MustParseInt64(stringStone[:digits/2])
			right := util.MustParseInt64(stringStone[digits/2:])

			updatedCount[left] += count
			updatedCount[right] += count
			totalStones += int64(count)
		}
		stoneCount = updatedCount
	}
	return totalStones
}

func parseStones(input string) []int {
	rawStones := strings.Split(input, " ")
	stones := make([]int, len(rawStones))
	for i, rs := range rawStones {
		stones[i] = util.MustParseInt(rs)
	}
	return stones
}
