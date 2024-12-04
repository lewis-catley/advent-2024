package main

import (
	"fmt"

	"github.com/lewis-catley/advent-2024/aoc"
	"github.com/lewis-catley/advent-2024/utils"
)

func main() {
	day1()

	fmt.Println("AOC 2024")
}

func day1() {
	l := aoc.Locations{}

	err := utils.ReadFile("files/1.txt", &l)
	if err != nil {
		panic(err)
	}

	l.Sort()
	fmt.Println("Day 1, Part 1 solution is: ", l.Difference())
	fmt.Println("Day 1, Part 2 solution is: ", l.SimilarityScore())
}

func day1_test() {
	l := aoc.NewLocations(
		[]int{3, 4, 2, 1, 3, 3},
		[]int{4, 3, 5, 3, 9, 3},
	)
	l.Sort()
	fmt.Println("Day 1, Part 1 solution is: ", l.Difference())
	fmt.Println("Day 1, Part 2 solution is: ", l.SimilarityScore())
}
