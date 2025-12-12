package main

import (
	"fmt"
	"strings"
)

func repeats(value string, prefix_len int) bool {
	return strings.Repeat(value[0:prefix_len], len(value)/prefix_len) == value
}

func day2() Solution {
	result1 := 0
	result2 := 0
	ids := ReadAndSplit(2, ",")
	for _, id := range ids {
		delim := strings.Index(id, "-")
		begin := StrToInt(id[0:delim])
		end := StrToInt(id[delim+1:])
		for i := begin; i <= end; i++ {
			value := fmt.Sprintf("%d", i)
			if len(value)%2 != 1 && repeats(value, len(value)/2) {
				result1 += i
			}
			for prefix_len := 1; prefix_len <= len(value)/2; prefix_len++ {
				if repeats(value, prefix_len) {
					result2 += i
					break
				}
			}
		}
	}
	return Solution{
		First:  fmt.Sprintf("%d", result1),
		Second: fmt.Sprintf("%d", result2),
	}
}
