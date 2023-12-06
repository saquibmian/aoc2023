package main

import (
	"strconv"
	"strings"
)

var day2part1 = RunFunc(func(input []string) string {
	type subset struct {
		red, green, blue int
	}
	type game struct {
		num     int
		subsets []subset
	}
	games := Map(input, func(line string) game {
		colon := strings.Index(line, ":")
		return game{
			num: MustInt(line[5:colon]),
			subsets: Map(strings.Split(line[colon+2:], ";"), func(subsetstr string) subset {
				var subset subset
				for _, cubes := range strings.Split(subsetstr, ",") {
					num, color, _ := strings.Cut(Trim(cubes), " ")
					switch Trim(color) {
					case "blue":
						subset.blue = MustInt(num)
					case "green":
						subset.green = MustInt(num)
					case "red":
						subset.red = MustInt(num)
					}
				}
				return subset
			}),
		}
	})
	var (
		red   = 12
		green = 13
		blue  = 14
		total int
	)
	for _, g := range games {
		match := All(g.subsets, func(s subset) bool {
			return s.red <= red && s.green <= green && s.blue <= blue
		})
		if match {
			total += g.num
		}
	}
	return strconv.Itoa(total)
})

var day2part2 = RunFunc(func(input []string) string {
	type subset struct {
		red, green, blue int
	}
	type game struct {
		num     int
		subsets []subset
	}
	games := Map(input, func(line string) game {
		colon := strings.Index(line, ":")
		return game{
			num: MustInt(line[5:colon]),
			subsets: Map(strings.Split(line[colon+2:], ";"), func(subsetstr string) subset {
				var subset subset
				for _, cubes := range strings.Split(subsetstr, ",") {
					num, color, _ := strings.Cut(Trim(cubes), " ")
					switch Trim(color) {
					case "blue":
						subset.blue = MustInt(num)
					case "green":
						subset.green = MustInt(num)
					case "red":
						subset.red = MustInt(num)
					}
				}
				return subset
			}),
		}
	})
	var powersum int
	for _, game := range games {
		var (
			fewestRed, fewestGreen, fewestBlue int
		)
		for _, ss := range game.subsets {
			fewestRed = Max(fewestRed, ss.red)
			fewestGreen = Max(fewestGreen, ss.green)
			fewestBlue = Max(fewestBlue, ss.blue)
		}
		powersum += (fewestRed * fewestGreen * fewestBlue)
	}
	return strconv.Itoa(powersum)
})
