package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Solution struct {
	First  string
	Second string
}

func solve(first int, second int) Solution {
	return Solution{
		First:  fmt.Sprintf("%d", first),
		Second: fmt.Sprintf("%d", second),
	}
}

func ReadAndSplit(day int, delim string) []string {
	path := fmt.Sprintf("inputs/input%d.txt", day)
	content, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return strings.Split(string(content), delim)
}

func ReadLines(day int) []string {
	return ReadAndSplit(day, "\n")
}

func StrToInt(str string) int {
	value, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return value
}

func mod(a, b int) int {
	m := a % b
	if a < 0 && b < 0 {
		m -= b
	}
	if a < 0 && b > 0 {
		m += b
	}
	return m
}
