package day01

import (
	"advent-of-code-go/util"
	"sort"
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

func (day Day01) Part2(data []string) string {
	var calories = []int{}
	for i, _ := range data {
		caloriesPerElve := 0
		for j := i; j < len(data); j++ {
			if data[j] == "" {
				break
			}
			caloriesPerElve += util.ParseToInt(data[j])
		}

		calories = append(calories, caloriesPerElve)
	}
	sort.Slice(calories, func(i, j int) bool {
		return calories[i] > calories[j]
	})
	totalOfTopThree := calories[0] + calories[1] + calories[2]
	return strconv.Itoa(totalOfTopThree)
}
