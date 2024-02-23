package util

import "strconv"

func GetDayAndPart(args []string) (int, int) {
	strDay := args[1]
	day := ParseToInt(strDay)
	if len(args) > 2 {
		part := ParseToInt(args[2])
		return day, part
	}

	return day, 1
}

func ParseToInt(in string) int {
	res, err := strconv.Atoi(in)
	if err != nil {
		panic(err)
	}
	return res
}
