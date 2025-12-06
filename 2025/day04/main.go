package main

import (
	"fmt"
	"os"
	"strings"
)

type Direction struct {
	dx int
	dy int
}

var (
	directions = []Direction{
		{0, -1},  // north
		{0, 1},   // south
		{1, 0},   // east
		{-1, 0},  // west
		{1, -1},  // north-east
		{1, 1},   // south-east
		{-1, -1}, // north-west
		{-1, 1},  // south-west
	}
)

func main() {
	bytes := Must(os.ReadFile("./input.txt"))
	lines := strings.Split(strings.TrimSpace(string(bytes)), "\n")

	grid := make([][]rune, len(lines))
	for i, line := range lines {
		grid[i] = []rune(line)
	}

	fmt.Println(Eval(grid))
}

func inBounds(grid [][]rune, x, y int) bool {
	return x >= 0 && y >= 0 && y < len(grid) && x < len(grid[0])
}

func Eval(grid [][]rune) int {
	result := 0

	for {
		removedThisPass := 0
		for y, row := range grid {
			for x := range row {
				if grid[y][x] == '.' {
					continue
				}
				count := 0
				for _, dir := range directions {
					newY := y + dir.dy
					newX := x + dir.dx
					if inBounds(grid, newX, newY) {
						if grid[newY][newX] == '@' {
							count++
						}
					}
				}

				if count < 4 {
					grid[y][x] = '.'
					removedThisPass++
				}
			}
		}
		result += removedThisPass
		if removedThisPass == 0 {
			break
		}
	}
	return result
}

func Must[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}
