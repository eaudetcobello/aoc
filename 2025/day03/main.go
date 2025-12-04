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
	stack := make([]int, 0, 12)
	budget := 3
	for _, num := range nums {
		var x int
		for len(stack) > 0 && budget > 0 && stack[len(stack)-1] < num {
			fmt.Printf("%d <= %d", stack[len(stack)-1], num)
			x, stack = stack[len(stack)-1], stack[:len(stack)-1]
			fmt.Printf(", pop %d ", x)
			fmt.Printf("(%v)\n", stack)
			budget--
		}
		stack = append(stack, num)
		fmt.Printf("push %d (%v)\n", num, stack)
	}

	fmt.Printf("final stack %v\n", stack)

	n := 0
	for _, d := range stack {
		n = n*10 + d
	}
	return n
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
