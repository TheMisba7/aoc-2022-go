package main

import (
	"advent-of-code-go/day"
	"advent-of-code-go/util"
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Day interface {
	Part1([]string) string
	Part2([]string) string
}

func loadData(day int) []string {
	filename := "day" + strconv.Itoa(day) + ".txt"
	readFile, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	out := []string{}
	for fileScanner.Scan() {
		out = append(out, fileScanner.Text())
	}

	readFile.Close()
	return out
}

func run(day Day, part int, data []string) string {
	if part == 1 {
		return day.Part1(data)
	}
	return day.Part2(data)
}

func main() {
	var days = make(map[int]Day)
	days[1] = day.Day01{}
	days[2] = day.Day02{}
	day, part := util.GetDayAndPart(os.Args)
	lines := loadData(day)
	var result = run(days[day], part, lines)

	fmt.Println(fmt.Sprintf("day %d, part %d => %s", day, part, result))

}
