package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

var day7part1 = RunFunc(func(input []string) string {
	type handtype int
	type hand struct {
		cards    []string
		bid      int
		handtype handtype
	}
	const (
		fiveOfAKind  handtype = 7
		fourOfAKind  handtype = 6
		fullHouse    handtype = 5
		threeOfAKind handtype = 4
		twoPair      handtype = 3
		onePair      handtype = 2
		highCard     handtype = 1
	)
	var hands []hand
	getHandType := func(cards []string) handtype {
		counts := map[string]int{
			"2": 0,
			"3": 0,
			"4": 0,
			"5": 0,
			"6": 0,
			"7": 0,
			"8": 0,
			"9": 0,
			"T": 0,
			"J": 0,
			"Q": 0,
			"K": 0,
			"A": 0,
		}
		for _, card := range cards {
			counts[card]++
		}
		for _, count := range counts {
			if count == 5 {
				return fiveOfAKind
			}
		}
		for _, count := range counts {
			if count == 4 {
				return fourOfAKind
			}
		}
		for i, count := range counts {
			if count == 3 {
				for j, count := range counts {
					if i == j {
						continue
					}
					if count == 2 {
						return fullHouse
					}
				}
				return threeOfAKind
			}
		}
		for i, count := range counts {
			if count == 2 {
				for j, count := range counts {
					if i == j {
						continue
					}
					if count == 2 {
						return twoPair
					}
				}
				return onePair
			}
		}
		return highCard
	}
	for _, line := range input {
		fields := strings.Fields(line)
		cards, bid := strings.Split(fields[0], ""), MustInt(fields[1])
		hands = append(hands, hand{
			cards:    cards,
			bid:      bid,
			handtype: getHandType(cards),
		})
	}
	getCardStrength := func(c string) int {
		return strings.Index("23456789TJQKA", c)
	}
	slices.SortFunc(hands, func(a, b hand) int {
		handtypeA, handtypeB := a.handtype, b.handtype
		if handtypeA > handtypeB {
			return 1
		}
		if handtypeA < handtypeB {
			return -1
		}
		for i := range a.cards {
			tiebreaker := getCardStrength(a.cards[i]) - getCardStrength(b.cards[i])
			if tiebreaker != 0 {
				return tiebreaker
			}
		}
		return 0
	})
	var total int
	for rank, hand := range hands {
		total += (hand.bid * (rank + 1))
	}
	return strconv.Itoa(total)
})

var day7part2 = RunFunc(func(input []string) string {
	type hand struct {
		cards []string
		bid   int
	}
	type handtype int
	handtypeString := func(h handtype) string {
		switch h {
		case 7:
			return "fiveOfAKind"
		case 6:
			return "fourOfAKind"
		case 5:
			return "fullHouse"
		case 4:
			return "threeOfAKind"
		case 3:
			return "twoPair"
		case 2:
			return "onePair"
		case 1:
			return "highCard"
		}
		panic("handtypestring")
	}
	const (
		fiveOfAKind  handtype = 7
		fourOfAKind  handtype = 6
		fullHouse    handtype = 5
		threeOfAKind handtype = 4
		twoPair      handtype = 3
		onePair      handtype = 2
		highCard     handtype = 1
	)
	var hands []hand
	for _, line := range input {
		fields := strings.Fields(line)
		cards, bid := strings.Split(fields[0], ""), MustInt(fields[1])
		hands = append(hands, hand{
			cards: cards,
			bid:   bid,
		})
	}
	getCardStrength := func(c string) int {
		return strings.Index("J23456789TQKA", c)
	}
	getHandType := func(h hand) handtype {
		counts := map[string]int{
			"2": 0,
			"3": 0,
			"4": 0,
			"5": 0,
			"6": 0,
			"7": 0,
			"8": 0,
			"9": 0,
			"T": 0,
			"J": 0,
			"Q": 0,
			"K": 0,
			"A": 0,
		}
		for _, card := range h.cards {
			counts[card]++
		}
		// This is important so that we don't consider J's count itself...
		// I needed help on this one.
		jCount := counts["J"]
		delete(counts, "J")
		for _, count := range counts {
			if count == 5 || count+jCount == 5 {
				return fiveOfAKind
			}
		}
		for _, count := range counts {
			if count == 4 || count+jCount == 4 {
				return fourOfAKind
			}
		}
		for i, count := range counts {
			if count == 3 || count+jCount == 3 {
				for j, count := range counts {
					if i == j {
						continue
					}
					if count == 2 {
						return fullHouse
					}
				}
				return threeOfAKind
			}
		}
		for i, count := range counts {
			if count == 2 || count+jCount == 2 {
				for j, count := range counts {
					if i == j {
						continue
					}
					if count == 2 {
						return twoPair
					}
				}
				return onePair
			}
		}
		return highCard
	}
	slices.SortFunc(hands, func(a, b hand) int {
		handtypeA, handtypeB := getHandType(a), getHandType(b)
		if handtypeA > handtypeB {
			return 1
		}
		if handtypeA < handtypeB {
			return -1
		}
		for i := range a.cards {
			tiebreaker := getCardStrength(a.cards[i]) - getCardStrength(b.cards[i])
			if tiebreaker != 0 {
				return tiebreaker
			}
		}
		return 0
	})
	for _, h := range hands {
		fmt.Printf("%s %s\n", h.cards, handtypeString(getHandType(h)))
	}
	var total int
	for rank, hand := range hands {
		total += (hand.bid * (rank + 1))
	}
	return strconv.Itoa(total)
})
