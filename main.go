package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

var days = [][]RunFunc{
	{
		day1part1,
		day1part2,
	},
	{
		day2part1,
		day2part2,
	},
	{
		day3part1,
		day3part2,
	},
	{
		day4part1,
		day4part2,
	},
	{
		day5part1,
		day5part2,
	},
	{
		day6part1,
		day6part2,
	},
}

type RunFunc func(input []string) string

func main() {
	if len(os.Args) != 2 {
		fmt.Fprint(os.Stderr, "usage: aoc2023 <day>\n")
		os.Exit(1)
	}
	dayArg := os.Args[1]
	day, err := strconv.Atoi(dayArg)
	if err != nil {
		fmt.Fprintln(os.Stderr, "not a number")
		os.Exit(1)
	}
	answer, err := answer(day)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println(answer)
}

func answer(day int) (string, error) {
	if day > len(days) {
		return "", errors.New("no answer for day")
	}
	parts := days[day-1]
	data, err := os.ReadFile(fmt.Sprintf("inputs/day%d.txt", day))
	if err != nil {
		return "", err
	}
	lines := Lines(string(data))
	run := parts[len(parts)-1]
	return run(lines), nil
}
