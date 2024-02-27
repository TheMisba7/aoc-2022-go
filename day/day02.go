package day

import (
	"strconv"
	"strings"
)

// Rock defeats Scissors, Scissors defeats Paper, and Paper defeats Rock
// rock: A:X
// paper: B:Y
// Scissors : c:Z
// x -> lose
// y -> draw
// z -> win
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
	totalScores := 0
	for _, line := range data {
		split := strings.Split(line, "")
		strategy := split[0] + " "
		if split[0] == "A" {
			switch split[2] {
			case "X":
				strategy += "Z"
			case "Y":
				strategy += "X"
			case "Z":
				strategy += "Y"
			}
		} else if split[0] == "C" {
			switch split[2] {
			case "X":
				strategy += "Y"
			case "Y":
				strategy += "Z"
			case "Z":
				strategy += "X"
			}
		} else {
			strategy = line
		}
		score := games[strategy]
		totalScores += score
	}
	return strconv.Itoa(totalScores)
}
