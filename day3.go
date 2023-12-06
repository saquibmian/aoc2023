package main

import (
	"strconv"
	"unicode"
)

var day3part1 = RunFunc(func(input []string) string {
	type number struct {
		num              int
		line, start, end int
	}
	var (
		symbolIndexes = Map(input, func(line string) []int {
			return IndexAll([]byte(line), func(b byte) bool {
				return b != '.' && !unicode.IsDigit(rune(b))
			})
		})
		numbers = MapIndex(input, func(line string, lineNumber int) []number {
			var numbers []number
			var start int = -1
			for end, c := range line {
				if unicode.IsDigit(c) {
					if start == -1 {
						start = end
					}
					continue
				}
				if start == -1 {
					continue
				}
				// we have a number we care about
				numbers = append(numbers, number{
					num:   MustInt(line[start:end]),
					line:  lineNumber,
					start: start,
					end:   end - 1,
				})
				start = -1
			}
			if start != -1 {
				numbers = append(numbers, number{
					num:   MustInt(line[start:]),
					line:  lineNumber,
					start: start,
					end:   len(line) - 1,
				})
			}
			return numbers
		})
		total int
	)
	for lineNumber := range input {
		for _, number := range numbers[lineNumber] {
			include := func(indexes []int) bool {
				return Any(indexes, func(i int) bool {
					return i <= number.end+1 && i >= number.start-1
				})
			}
			if lineNumber != 0 && include(symbolIndexes[lineNumber-1]) {
				total += number.num
			} else if include(symbolIndexes[lineNumber]) {
				total += number.num
			} else if lineNumber != len(input)-1 && include(symbolIndexes[lineNumber+1]) {
				total += number.num
			}
		}
	}
	return strconv.Itoa(total)
})

var day3part2 = RunFunc(func(input []string) string {
	type number struct {
		num              int
		line, start, end int
	}
	var (
		numbers = MapIndex(input, func(line string, lineNumber int) []number {
			var numbers []number
			var start int = -1
			for end, c := range line {
				if unicode.IsDigit(c) {
					if start == -1 {
						start = end
					}
					continue
				}
				if start == -1 {
					continue
				}
				// we have a number we care about
				numbers = append(numbers, number{
					num:   MustInt(line[start:end]),
					line:  lineNumber,
					start: start,
					end:   end - 1,
				})
				start = -1
			}
			if start != -1 {
				numbers = append(numbers, number{
					num:   MustInt(line[start:]),
					line:  lineNumber,
					start: start,
					end:   len(line) - 1,
				})
			}
			return numbers
		})
		total int
	)
	for lineNumber, line := range input {
		for symbolIndex, symbol := range line {
			if symbol != rune('*') {
				continue
			}
			var (
				above, curr, below []number
			)
			if lineNumber > 0 {
				above = FindAll(numbers[lineNumber-1], func(n number) bool {
					return n.start-1 <= symbolIndex && symbolIndex <= n.end+1
				})
			}
			curr = FindAll(numbers[lineNumber], func(n number) bool {
				return n.start-1 <= symbolIndex && symbolIndex <= n.end+1
			})
			if lineNumber < len(input)-1 {
				below = FindAll(numbers[lineNumber+1], func(n number) bool {
					return n.start-1 <= symbolIndex && symbolIndex <= n.end+1
				})
			}
			if (len(above) + len(curr) + len(below)) > 2 {
				continue
			}
			if len(above) == 2 {
				// above and above
				total += (above[0].num * above[1].num)
			} else if len(above) == 1 && len(curr) == 1 {
				// above and same
				total += (above[0].num * curr[0].num)
			} else if len(curr) == 2 {
				// same and same
				total += (curr[0].num * curr[1].num)
			} else if len(curr) == 1 && len(below) == 1 {
				// same and below
				total += (curr[0].num * below[0].num)
			} else if len(below) == 2 {
				// below and below
				total += (below[0].num * below[1].num)
			} else if len(above) == 1 && len(below) == 1 {
				total += (above[0].num * below[0].num)
			}
		}
	}
	return strconv.Itoa(total)
})
