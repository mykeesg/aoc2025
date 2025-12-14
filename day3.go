package main

func day3solution(length int, line string) int {
	positions := make([]int, length)
	lastPos := -1
	for idx := range len(positions) {
		positions[idx] = -1
		for jdx := len(line) - length + idx; jdx > lastPos; jdx -= 1 {
			if positions[idx] == -1 || line[jdx] >= line[positions[idx]] {
				positions[idx] = jdx
			}
		}
		lastPos = positions[idx]
	}
	sum := 0
	for idx := range len(positions) {
		value := int(line[positions[idx]] - '0')
		sum = 10*sum + value
	}
	return sum
}

func day3() Solution {
	first := 0
	second := 0
	for _, line := range ReadLines(3) {
		first += day3solution(2, line)
		second += day3solution(12, line)
	}
	return solve(first, second)
}
