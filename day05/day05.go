package day05

import (
	"slices"
	"strings"

	"github.com/ruegerj/aoc-2024/pkg/common"
	"github.com/ruegerj/aoc-2024/pkg/util"
)

type Day05 struct{}

func (d Day05) Part1(input string) *common.Solution {
	inputParts := strings.Split(input, "\n\n")
	rules := parseRules(inputParts[0])
	manuals := parseManuals(inputParts[1])

	middlePageSum := 0

	for _, manual := range manuals {
		isValid := isValidManual(manual, rules)

		if isValid {
			middlePageSum += getMiddlePage(manual)
		}
	}

	return common.NewSolution(1, middlePageSum)
}

func (d Day05) Part2(input string) *common.Solution {
	inputParts := strings.Split(input, "\n\n")
	rules := parseRules(inputParts[0])
	manuals := parseManuals(inputParts[1])

	middlePageSum := 0

	for _, manual := range manuals {
		isValid := isValidManual(manual, rules)

		if isValid {
			continue
		}

		sortManual(manual, rules)
		middlePageSum += getMiddlePage(manual)
	}

	return common.NewSolution(2, middlePageSum)
}

type rule struct {
	page   int
	before int
}

func sortManual(manual []int, rules []rule) {
	slices.SortFunc(manual, func(a, b int) int {
		for _, r := range rules {
			if a == r.page && b == r.before {
				return -1
			}

			if a == r.before && b == r.page {
				return 1
			}
		}

		return 0
	})
}

func isValidManual(manual []int, rules []rule) bool {
	valid := true
	printedPages := make([]int, 0)
	for _, page := range manual {
		printedPages = append(printedPages, page)

		for _, rule := range rules {
			if page != rule.page {
				continue
			}

			valid = !slices.Contains(printedPages, rule.before)

			if !valid {
				break
			}
		}

		if !valid {
			break
		}
	}

	return valid
}

func getMiddlePage(manual []int) int {
	return manual[len(manual)/2]
}

func parseRules(input string) []rule {
	lines := util.Lines(input)
	rules := make([]rule, len(lines))

	for i, line := range lines {
		parts := strings.Split(line, "|")
		rules[i] = rule{
			page:   util.MustParseInt(parts[0]),
			before: util.MustParseInt(parts[1]),
		}
	}

	return rules
}

func parseManuals(input string) [][]int {
	lines := util.Lines(input)
	manuals := make([][]int, len(lines))

	for i, line := range lines {
		pages := strings.Split(line, ",")
		manuals[i] = make([]int, len(pages))
		for j, page := range pages {
			manuals[i][j] = util.MustParseInt(page)
		}
	}

	return manuals
}
