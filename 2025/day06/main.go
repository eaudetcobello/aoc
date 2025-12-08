package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type Operator string

const (
	OpPlus Operator = "+"
	OpMult Operator = "*"
)

func parseOperator(s string) (Operator, error) {
	opMap := map[string]Operator{
		"+": OpPlus,
		"*": OpMult,
	}

	op, ok := opMap[s]
	if !ok {
		return "", fmt.Errorf("unknown error: %q", s)
	}
	return op, nil
}

func main() {
	bytes := Must(os.ReadFile("./input.txt"))
	lines := strings.Split(strings.TrimSpace(string(bytes)), "\n")
	Eval(lines)
}

func Eval(lines []string) [][]int {
	maxLen := maxLength(lines)
	paddedLines := pad(lines, maxLen)
	cols := cols(paddedLines)
	groupedCols := groupCols(cols)
	problems := parseProblems(groupedCols)
	total := calc(problems)

	fmt.Println(total)

	return nil
}

// Problem is a list of integers that an operator must act on
type Problem struct {
	elems []int
	op    Operator
}

func calc(problems []Problem) int {
	var acc int
	for _, prob := range problems {
		var local int
		switch prob.op {
		case OpMult:
			local = 1
			for _, elem := range prob.elems {
				local *= elem
			}
		case OpPlus:
			for _, elem := range prob.elems {
				local += elem
			}
		}
		acc += local
	}

	return acc
}

func parseProblems(groups [][]string) []Problem {
	out := make([]Problem, len(groups))
	for i := range groups {
		out[i] = parseProblem(groups[i])
	}
	return out
}

// parseProblem parses a group of columns into a
// Problem struct
func parseProblem(cols []string) Problem {
	var longestCol int
	for _, col := range cols {
		var count int
		for _, r := range col {
			if unicode.IsDigit(r) {
				count++
			}
		}
		if count > longestCol {
			longestCol = count
		}
	}

	numbers := make([]int, len(cols))

	for i := range cols {
		var number strings.Builder
		for j := range longestCol {
			if cols[i][j] != ' ' {
				number.WriteByte(cols[i][j])
			}
		}
		numbers[i] = Must(strconv.Atoi(number.String()))
	}

	op, _ := parseOperator(string(cols[0][len(cols[0])-1]))
	prob := Problem{
		elems: numbers,
		op:    op,
	}

	return prob
}

// accumulates columns until it encouters a separator (all-whitespace column),
// then forms a group, finally resets accumulator
// in the returned slice, each element is a slice of
// the columns of a problem.
func groupCols(cols []string) [][]string {
	var out [][]string
	var acc []string
	for j := range len(cols) {
		if strings.TrimSpace(cols[j]) == "" {
			if len(acc) > 0 {
				out = append(out, acc)
				acc = make([]string, 0)
			}
			continue
		}
		acc = append(acc, cols[j])
	}
	out = append(out, acc)
	return out
}

// cols returns cols from a string slice where
// all rows have same length
func cols(lines []string) []string {
	cols := make([]string, len(lines[0])) // as many columns as the line length

	for j := range len(lines[0]) {
		var sb strings.Builder
		for i := range lines {
			sb.WriteByte(lines[i][j])
		}
		cols[j] = sb.String()
	}

	return cols
}

func pad(lines []string, maxLen int) []string {
	out := make([]string, len(lines))
	for i := range lines {
		out[i] = fmt.Sprintf("%-*s", maxLen, lines[i])
	}
	return out
}

func maxLength(lines []string) int {
	maxLen := 0

	for i := range lines {
		if len(lines[i]) > maxLen {
			maxLen = len(lines[i])
		}
	}

	return maxLen
}

func Must[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}
