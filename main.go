package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/ruegerj/aoc-2024/day01"
	"github.com/ruegerj/aoc-2024/day02"
	"github.com/ruegerj/aoc-2024/pkg/common"
	"github.com/ruegerj/aoc-2024/pkg/util"
)

func main() {
	fmt.Println(`     ___       ______     ______     ___     ___    ___    _  _    `)
	fmt.Println(`    /   \     /  __  \   /      |   |__ \   / _ \  |__ \  | || |   `)
	fmt.Println(`   /  ^  \   |  |  |  | |  ,----'      ) | | | | |    ) | | || |_  `)
	fmt.Println(`  /  /_\  \  |  |  |  | |  |          / /  | | | |   / /  |__   _| `)
	fmt.Println(" /  _____  \\ |  `--'  | |  `----.    / /_  | |_| |  / /_     | |   ")
	fmt.Println(`/__/     \__\ \______/   \______|   |____|  \___/  |____|    |_|   `)
	fmt.Println("-------------------------------------------------------------------")
	fmt.Println("🎄 Happy Coding & festive season")

	dayArg := os.Args[1]
	printHelp := strings.Contains(dayArg, "help") || strings.Contains(dayArg, "h")

	if printHelp {
		fmt.Println("usage: go run . <day-nr>")
		return
	}

	dayNr, err := strconv.Atoi(dayArg)

	if err != nil {
		fmt.Println("❌ Invalid day number...")
		return
	}

	dayRegistry := map[int]common.Day{
		1: day01.Day01{},
		2: day02.Day02{},
	}
	requestedDay := dayRegistry[dayNr]

	if requestedDay == nil {
		fmt.Println("🛠  Not implemented")
		return
	}

	runDay(dayNr, requestedDay)
}

func runDay(nr int, day common.Day) {
	input := common.LoadDailyInput(nr)
	normalizedNr := util.PadNumber(nr)

	fmt.Printf("\n⭐️ Day %s\n", normalizedNr)

	start1 := time.Now()
	solution1 := day.Part1(input)
	solution1.Print(time.Since(start1))

	start2 := time.Now()
	solution2 := day.Part2(input)
	solution2.Print(time.Since(start2))
}