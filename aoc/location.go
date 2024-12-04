package aoc

import (
	"sort"
	"strconv"
	"strings"

	"github.com/lewis-catley/advent-2024/utils"
)

type Locations struct {
	x []int
	y []int
}

func NewLocations(x, y []int) Locations {
	return Locations{
		x: x,
		y: y,
	}
}

func (l *Locations) Sort() {
	sort.Ints(l.x)
	sort.Ints(l.y)
}

func (l Locations) SimilarityScore() int {
	out := 0
	occurance := utils.CountOccurnace(l.y)
	for _, i := range l.x {
		if count, ok := occurance[i]; ok {
			out += count * i
		}
	}
	return out
}

func (l Locations) Difference() uint {
	var out uint = 0
	for i := range l.x {
		diff := (l.x[i] - l.y[i])
		if diff < 0 {
			diff = diff * -1
		}
		out += uint(diff)
	}
	return out
}

func (l *Locations) Parse(b []byte) error {
	lines := strings.Split(string(b), "\n")
	// Remove empty new line
	lines = lines[:len(lines)-1]

	totVals := len(lines)

	x := make([]int, totVals)
	y := make([]int, totVals)

	for i, line := range lines {
		vals := strings.Split(line, " ")
		xVal, _ := strconv.Atoi(vals[0])
		yVal, _ := strconv.Atoi(vals[len(vals)-1])

		x[i] = xVal
		y[i] = yVal

	}
	l.x = x
	l.y = y
	return nil
}
