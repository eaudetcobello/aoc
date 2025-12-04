package main

import (
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"
)

func main() {
	bytes := Must(os.ReadFile("./input.txt"))
	s := string(bytes)
	s = strings.TrimRight(s, "\r\n")
	lines := strings.SplitSeq(s, "\n")

	sum := big.NewInt(0)
	for line := range lines {
		line = strings.TrimRight(line, "\r\n")
		nums := Map(strings.Split(line, ""), func(s string) int {
			return Must(strconv.Atoi(s))
		})
		sum.Add(sum, Eval(nums))
	}
	fmt.Println(sum)
}

func Eval(nums []int) *big.Int {
	stack := make([]int, 0, 12)
	budget := len(nums) - 12
	for _, num := range nums {
		for len(stack) > 0 && budget > 0 && stack[len(stack)-1] < num {
			_, stack = stack[len(stack)-1], stack[:len(stack)-1]
			budget--
		}
		stack = append(stack, num)
	}

	stack = stack[:len(stack)-budget]

	n := big.NewInt(0)
	ten := big.NewInt(10)
	for _, d := range stack {
		n.Mul(n, ten)
		n.Add(n, big.NewInt(int64(d)))
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
