package day

import (
	"strconv"
)

var games = map[string]int{
	"A X": 4,
	"A Y": 8,
	"A Z": 3,
	"B X": 1,
	"B Y": 5,
	"B Z": 9,
	"C X": 7,
	"C Y": 2,
	"C Z": 6}

type Day02 struct{}

func (day Day02) Part1(data []string) string {
	totalScores := 0
	for _, key := range data {
		score := games[key]
		totalScores += score
	}
	return strconv.Itoa(totalScores)
}

func (day Day02) Part2(data []string) string {
	return ""
}
