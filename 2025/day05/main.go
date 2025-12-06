package main

import (
	"cmp"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	bytes := Must(os.ReadFile("./input.txt"))
	parts := strings.Split(strings.TrimSpace(string(bytes)), "\n\n")

	rangeStrs := strings.Split(parts[0], "\n")
	ingredientStrs := strings.Split(parts[1], "\n")

	ranges := make([][]int, len(rangeStrs))
	for i, line := range rangeStrs {
		parts := strings.Split(line, "-")
		start, end := Must(strconv.Atoi(parts[0])), Must(strconv.Atoi(parts[1]))
		ranges[i] = []int{start, end}
	}

	ingredients := make([]int, len(ingredientStrs))
	for i, ingredient := range ingredientStrs {
		ingredients[i] = Must(strconv.Atoi(ingredient))
	}

	fmt.Println(CountFresh(ingredients, Merge(ranges)))
}

func CountFresh(ingredients []int, ranges [][]int) int {
	total := 0

	searchFunc := func(bucket []int, a int) int {
		if a < bucket[0] {
			return 1
		} else if a > bucket[1] {
			return -1
		}
		return 0
	}

	for _, ingredient := range ingredients {
		_, found := slices.BinarySearchFunc(ranges, ingredient, searchFunc)
		if found {
			total++
		}
	}

	return total
}

// Merge overlapping ranges in a 2d slice
// Example: [[1,5],[3,6]] => [[1,6]]
func Merge(ranges [][]int) [][]int {
	rangeCmp := func(a, b []int) int {
		return cmp.Compare(a[0], b[0])
	}

	slices.SortFunc(ranges, rangeCmp)

	current := ranges[0]
	merged := make([][]int, 0, len(ranges))
	for i := 1; i < len(ranges); i++ {
		if ranges[i][0] <= current[1] {
			maxEnd := math.Max(f64(ranges[i][1]), f64(current[1]))
			current = []int{current[0], int(maxEnd)}
		} else {
			merged = append(merged, current)
			current = ranges[i]
		}
	}

	merged = append(merged, current)

	return merged
}

func f64(n int) float64 {
	return float64(n)
}

func Must[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}
