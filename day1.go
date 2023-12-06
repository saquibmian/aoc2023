package main

import (
	"strconv"
	"strings"
)

var day1part1 = RunFunc(func(input []string) string {
	var runningTotal int
	for _, line := range input {
		var (
			firstDigit string
			lastDigit  string
		)
		for _, char := range line {
			if _, err := strconv.Atoi(string(char)); err != nil {
				continue
			}
			if firstDigit == "" {
				firstDigit = string(char)
				lastDigit = string(char)
			} else {
				lastDigit = string(char)
			}
		}
		runningTotal += MustInt(firstDigit + lastDigit)
	}
	return strconv.Itoa(runningTotal)
})

var day1part2 = RunFunc(func(input []string) string {
	validInts := map[string]int{
		"1":     1,
		"one":   1,
		"2":     2,
		"two":   2,
		"3":     3,
		"three": 3,
		"4":     4,
		"four":  4,
		"5":     5,
		"five":  5,
		"6":     6,
		"six":   6,
		"7":     7,
		"seven": 7,
		"8":     8,
		"eight": 8,
		"9":     9,
		"nine":  9,
	}
	var runningTotal int
	for _, line := range input {
		var (
			firstDigit      int
			firstDigitIndex = -1
			lastDigit       int
			lastDigitIndex  = -1
		)
		for str, val := range validInts {
			if index := strings.Index(line, str); index != -1 {
				if firstDigitIndex == -1 {
					firstDigit, firstDigitIndex = val, index
					lastDigit, lastDigitIndex = val, index
				} else if firstDigitIndex > index {
					firstDigit, firstDigitIndex = val, index
				}
			}
			if index := strings.LastIndex(line, str); index != -1 {
				if lastDigitIndex < index {
					lastDigit, lastDigitIndex = val, index
				}
			}
		}
		runningTotal += (firstDigit * 10) + lastDigit
	}
	return strconv.Itoa(runningTotal)
})
