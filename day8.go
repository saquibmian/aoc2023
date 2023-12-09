package main

import (
	"regexp"
	"strconv"
	"strings"
)

var day8part1 = RunFunc(func(input []string) string {
	type node struct {
		name        string
		left, right string
	}
	nodes := map[string]node{}
	instructions := strings.Split(input[0], "")
	for _, line := range input[2:] {
		r := regexp.MustCompile(`[A-Z]{3}`)
		matched := r.FindAllString(line, -1)
		nodes[matched[0]] = node{
			name:  matched[0],
			left:  matched[1],
			right: matched[2],
		}
	}
	var (
		moves int
		start = nodes["AAA"]
	)
	for {
		for _, instruction := range instructions {
			if start.name == "ZZZ" {
				break
			}
			moves++
			switch instruction {
			case "L":
				start = nodes[start.left]
			case "R":
				start = nodes[start.right]
			default:
				panic(instruction)
			}
		}
		if start.name == "ZZZ" {
			break
		}
	}
	return strconv.Itoa(moves)
})

var day8part2 = RunFunc(func(input []string) string {
	type node struct {
		name        string
		left, right string
	}
	nodes := map[string]node{}
	instructions := strings.Split(input[0], "")
	for _, line := range input[2:] {
		r := regexp.MustCompile(`[A-Z0-9]{3}`)
		matched := r.FindAllString(line, -1)
		nodes[matched[0]] = node{
			name:  matched[0],
			left:  matched[1],
			right: matched[2],
		}
	}
	var (
		moves               int
		movesUntilSatisfied []int
		starts              []node
	)
	for _, node := range nodes {
		if strings.HasSuffix(node.name, "A") {
			movesUntilSatisfied = append(movesUntilSatisfied, 0)
			starts = append(starts, node)
		}
	}
	for {
		for _, instruction := range instructions {
			for I, node := range starts {
				if strings.HasSuffix(node.name, "Z") {
					movesUntilSatisfied[I] = moves
				}
			}
			if All(movesUntilSatisfied, func(i int) bool { return i != 0 }) {
				break
			}
			if All(starts, func(n node) bool { return strings.HasSuffix(n.name, "Z") }) {
				break
			}
			moves++
			switch instruction {
			case "L":
				for i := range starts {
					starts[i] = nodes[starts[i].left]
				}
			case "R":
				for i := range starts {
					starts[i] = nodes[starts[i].right]
				}
			default:
				panic(instruction)
			}
		}
		if All(movesUntilSatisfied, func(i int) bool { return i != 0 }) {
			break
		}
		if All(starts, func(n node) bool { return strings.HasSuffix(n.name, "Z") }) {
			break
		}
	}
	return strconv.Itoa(LCM(movesUntilSatisfied[0], movesUntilSatisfied[1], movesUntilSatisfied[2:]...))
})
