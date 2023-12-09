package main

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDay1(t *testing.T) {
	testDayPart(t, 1, 1, "142")
	testDayPart(t, 1, 2, "281")
}

func TestDay2(t *testing.T) {
	testDayPart(t, 2, 1, "8")
	testDayPart(t, 2, 2, "2286")
}

func TestDay3(t *testing.T) {
	testDayPart(t, 3, 1, "4361")   // on full input 521515
	testDayPart(t, 3, 2, "467835") // on full input 69527306
}

func TestDay4(t *testing.T) {
	testDayPart(t, 4, 1, "13") // 25010
	testDayPart(t, 4, 2, "30") // 9924412
}

func TestDay5(t *testing.T) {
	testDayPart(t, 5, 1, "35") // 177942185
	testDayPart(t, 5, 2, "46") // 69841803
}

func TestDay6(t *testing.T) {
	testDayPart(t, 6, 1, "288")   // 1312850
	testDayPart(t, 6, 2, "71503") // 36749103
}

func TestDay7(t *testing.T) {
	testDayPart(t, 7, 1, "6440") // 256448566
	testDayPart(t, 7, 2, "5905") // 254412181
}

func TestDay8(t *testing.T) {
	testDayPart(t, 8, 1, "2") // 16409
	testDayPart(t, 8, 2, "6") // 11795205644011
}

func testDayPart(t *testing.T, day int, part int, expected string) {
	sampleData, err := os.ReadFile(fmt.Sprintf("samples/day%dpart%d.txt", day, part))
	require.NoError(t, err)
	lines := strings.Split(string(sampleData), "\n")
	sampleAnswer := days[day-1][part-1](lines)
	require.Equal(t, expected, sampleAnswer)
}
