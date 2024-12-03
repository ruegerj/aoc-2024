package util

import (
	"regexp"
)

func MatchNamedSubgroups(exp *regexp.Regexp, input string) map[string]string {
	match := exp.FindStringSubmatch(input)

	subgroupValues := make(map[string]string)
	for i, name := range exp.SubexpNames() {
		if i > 0 && i <= len(match) {
			subgroupValues[name] = match[i]
		}
	}

	return subgroupValues
}

func MatchAllNamedSubgroups(exp *regexp.Regexp, input string) []map[string]string {
	matches := exp.FindAllStringSubmatch(input, -1)

	subgroupValues := make([]map[string]string, len(matches))
	for i, match := range matches {
		subgroupValues[i] = make(map[string]string)
		for y, name := range exp.SubexpNames() {
			if y > 0 && y <= len(match) {
				subgroupValues[i][name] = match[y]
			}
		}
	}

	return subgroupValues
}
