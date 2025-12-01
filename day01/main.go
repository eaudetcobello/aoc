package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)


func floorDiv100(a int) int {
	if a >= 0 {
		return a / 100
	}
	return (a - 99) / 100
}

func turnWrap(pos int, n int) (wraps, newPos int) {
	start := pos
	end := pos + n
	if n > 0 {
		wraps = (end-1)/100 - start/100
	} else if n < 0 {
		wraps = floorDiv100(start-1) - floorDiv100(end)
	} else {
		wraps = 0
	}
	newPos = ((end % 100) + 100) % 100
	return
}

func dirStrToInt(s string) int {
	dir := string(s[0])
	n, _ := strconv.Atoi(s[1:])
	switch dir {
	case "L":
		return -n
	default:
		return n
	}
}

func main() {
	bytes, _ := os.ReadFile("./input.txt")
	s := string(bytes)
	s = strings.TrimRight(s, "\r\n")
	lines := strings.Split(s, "\n")

	c, d := turnWrap(50, 1000)
	fmt.Printf("%d,%d\n", c, d)


	pos := 50
	var landedZeroN int
	var passedZeroN int
	for _, line := range lines {
		a, b := turnWrap(pos, dirStrToInt(line))
		pos = b
		if pos == 0 {
			landedZeroN++
		}
		passedZeroN += int(math.Abs(float64(a)))
	}


	fmt.Printf("%d,%d\n", landedZeroN + passedZeroN, pos)
}
