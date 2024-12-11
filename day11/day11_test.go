package day11

import (
	"fmt"
	"os"
	"path"
	"testing"
)

var dailyInput string

func TestMain(m *testing.M) {
	input, err := os.ReadFile(path.Join("..", "data", "11.txt"))

	if err != nil {
		fmt.Println("Failed to load input file", err)
		os.Exit(1)
	}

	dailyInput = string(input)

	code := m.Run()
	os.Exit(code)
}

func TestPart1(t *testing.T) {
	expected := int64(202019)
	solution := Day11{}.Part1(dailyInput)

	if solution.Result.(int64) != expected {
		t.Errorf("Expected %d, produced %d", expected, solution.Result)
	}
}

func TestPart2(t *testing.T) {
	expected := int64(239321955280205)
	solution := Day11{}.Part2(dailyInput)

	if solution.Result.(int64) != expected {
		t.Errorf("Expected %d, produced %d", expected, solution.Result)
	}
}
