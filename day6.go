package main

import (
	"strconv"
	"strings"
)

var day6part1 = RunFunc(func(input []string) string {
	type race struct {
		totalMs          int
		recordDistanceMm int
	}
	var (
		totalMss = Map(
			strings.Fields(strings.TrimPrefix(input[0], "Time:")),
			func(s string) int { return MustInt(s) },
		)
		recordDistanceMms = Map(
			strings.Fields(strings.TrimPrefix(input[1], "Distance:")),
			func(s string) int { return MustInt(s) },
		)
		races []race
	)
	for i := range totalMss {
		races = append(races, race{totalMs: totalMss[i], recordDistanceMm: recordDistanceMms[i]})
	}
	var total int
	for _, race := range races {
		var wins int
		for i := 1; i < race.totalMs; i++ {
			if i*(race.totalMs-i) > race.recordDistanceMm {
				wins++
			}
		}
		if total == 0 {
			total = wins
		} else {
			total *= wins
		}
	}
	return strconv.Itoa(total)
})

var day6part2 = RunFunc(func(input []string) string {
	type race struct {
		totalMs          int
		recordDistanceMm int
	}
	r := race{
		totalMs:          MustInt(strings.ReplaceAll(strings.TrimPrefix(input[0], "Time:"), " ", "")),
		recordDistanceMm: MustInt(strings.ReplaceAll(strings.TrimPrefix(input[1], "Distance:"), " ", "")),
	}
	var wins int
	for i := 1; i < r.totalMs; i++ {
		if i*(r.totalMs-i) > r.recordDistanceMm {
			wins++
		}
	}
	return strconv.Itoa(wins)
})
