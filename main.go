package main

import "fmt"

func main() {
	solutions := make([]Solution, 0, 25)
	// solutions = append(solutions, day1())
	// solutions = append(solutions, day2())
	solutions = append(solutions, day3())

	for idx, sol := range solutions {
		fmt.Printf("Day %d\n", (idx + 1))
		fmt.Printf("⭐🌑   1st :: %s\n", sol.First)
		fmt.Printf("⭐⭐   2nd :: %s\n", sol.Second)
	}
}
