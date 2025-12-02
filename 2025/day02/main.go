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
	lines := strings.SplitSeq(s, ",")

	total := 0
	for line := range lines {
		parts := strings.Split(line, "-")
		start := Must(strconv.Atoi(parts[0]))
		end := Must(strconv.Atoi(parts[1]))
		nums := Filter(Range(start, end), func(s string) bool {
			return !IsValid(s)
		})
		for _, i := range nums {
			total += Must(strconv.Atoi(i))
		}
	}

	fmt.Println(total)
}

func Must[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}

func Range(start, end int) []string {
	result := make([]string, 0, end-start+1)
	for i := start; i <= end; i++ {
		result = append(result, strconv.Itoa(i))
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

func IsValid(s string) bool {
	if len(s)%2 == 0 {
		mid := len(s) / 2
		left, right := s[:mid], s[mid:]
		if strings.EqualFold(left, right) {
			return false
		}
	}
	return true
}
