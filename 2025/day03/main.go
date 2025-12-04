package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	bytes := Must(os.ReadFile("./input.txt"))
	s := string(bytes)
	s = strings.TrimRight(s, "\r\n")
	lines := strings.SplitSeq(s, "\n")

	sum := 0
	for line := range lines {
		line = strings.TrimRight(line, "\r\n")
		nums := Map(strings.Split(line, ""), func(s string) int {
			return Must(strconv.Atoi(s))
		})
		sum += Eval(nums)
	}
	fmt.Println(sum)
}

func Eval(nums []int) int {
	bestTens := 0
	bestTwoDigit := 0
	for _, num := range nums {
		if (bestTens*10)+num > bestTwoDigit {
			bestTwoDigit = (bestTens * 10) + num
		}
		if num > bestTens {
			bestTens = num
		}
	}
	return bestTwoDigit
}

func Must[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}

func Map[T any, E any](slice []T, mapFunc func(T) E) []E {
	result := make([]E, 0, len(slice))
	for _, elem := range slice {
		result = append(result, mapFunc(elem))
	}
	return result
}

func Filter[T any](slice []T, predicate func(T) bool) []T {
	result := make([]T, 0, len(slice))
	for _, elem := range slice {
		if predicate(elem) {
			result = append(result, elem)
		}
	}
	return result
}
