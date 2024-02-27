package day

import (
	"fmt"
	"strconv"
	"strings"
)

type Day03 struct{}

var priorityMap = getPriority()

func getPriority() map[string]int {
	var priority = map[string]int{}
	lower := 'a'
	lowerPriority := 1
	upper := 'A'
	upperPriority := 27
	for lower <= 'z' {
		priority[fmt.Sprintf("%c", lower)] = lowerPriority
		priority[fmt.Sprintf("%c", upper)] = upperPriority
		upperPriority++
		lowerPriority++
		lower++
		upper++
	}
	return priority
}
func (day Day03) Part1(data []string) string {
	var totalPriorities = 0
	for _, line := range data {
		half := len(line) / 2
		first := line[:half]
		second := line[half:]
		for i := 0; i < len(first); i++ {
			var str string = fmt.Sprintf("%c", first[i])
			if strings.Contains(second, str) && !strings.Contains(first[:i], str) {
				totalPriorities += priorityMap[str]
			}
		}
	}
	return strconv.Itoa(totalPriorities)
}

func (day Day03) Part2(data []string) string {
	return ""
}
