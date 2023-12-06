package main

import (
	"math"
	"slices"
	"strconv"
	"strings"
)

var day4part1 = RunFunc(func(input []string) string {
	type card struct {
		winning []int
		hand    []int
	}
	var cards []card
	for _, line := range input {
		line := line[strings.Index(line, ":")+1:]
		c := card{}
		winning, hand, _ := strings.Cut(line, " | ")

		for _, w := range strings.Split(winning, " ") {
			if Trim(w) == "" {
				continue
			}
			c.winning = append(c.winning, MustInt(Trim(w)))
		}
		for _, w := range strings.Split(hand, " ") {
			if Trim(w) == "" {
				continue
			}
			c.hand = append(c.hand, MustInt(Trim(w)))
		}
		cards = append(cards, c)
	}
	var total int
	for _, card := range cards {
		numWinning := len(Filter(card.hand, func(i int) bool {
			return slices.Index(card.winning, i) != -1
		}))
		total += int(math.Pow(2, float64(numWinning-1)))
	}
	return strconv.Itoa(total)
})

var day4part2 = RunFunc(func(input []string) string {
	type card struct {
		number  int
		winning []int
		hand    []int
	}
	var cards []card
	for i, line := range input {
		line := line[strings.Index(line, ":")+1:]
		c := card{number: i + 1}
		winning, hand, _ := strings.Cut(line, " | ")

		for _, w := range strings.Split(winning, " ") {
			if Trim(w) == "" {
				continue
			}
			c.winning = append(c.winning, MustInt(Trim(w)))
		}
		for _, w := range strings.Split(hand, " ") {
			if Trim(w) == "" {
				continue
			}
			c.hand = append(c.hand, MustInt(Trim(w)))
		}
		cards = append(cards, c)
	}
	var (
		toProcess      = append([]card{}, cards...)
		totalProcessed int
	)
	for len(toProcess) != 0 {
		totalProcessed++
		currentCard := toProcess[0]
		toProcess = toProcess[1:]
		if currentCard.number < len(cards) {
			numWinning := len(Filter(currentCard.hand, func(i int) bool {
				return slices.Index(currentCard.winning, i) != -1
			}))
			if numWinning > 0 {
				toAdd := cards[currentCard.number:Min(len(cards), currentCard.number+numWinning)]
				toProcess = append(toProcess, toAdd...)
			}
		}
	}
	return strconv.Itoa(totalProcessed)
})
