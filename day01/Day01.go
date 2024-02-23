package day01

import (
	"advent-of-code-go/util"
	"strconv"
)

type Day01 struct{}

func (day Day01) Part1(data []string) string {
	var totalCalories int = 0
	for i, _ := range data {
		caloriesPerElve := 0
		for j := i; j < len(data); j++ {
			if data[j] == "" {
				break
			}
			caloriesPerElve += util.ParseToInt(data[j])
		}

		if caloriesPerElve > totalCalories {
			totalCalories = caloriesPerElve
		}
	}
	return strconv.Itoa(totalCalories)
}

func (day Day01) Part2([]string) string {
	return ""
}
