package main

func day1() Solution {
	stops := 0
	passes := 0
	position := 50

	for _, line := range ReadLines(1) {
		if len(line) == 0 {
			break // new line at end
		}

		rotation := StrToInt(line[1:])
		delta := mod(rotation, 100)
		passes += rotation / 100

		if line[0] == 'R' {
			position += delta
			if position >= 100 {
				passes += 1
				position -= 100
			}
		} else {
			old_pos := position
			position -= delta
			if position < 0 {
				if old_pos != 0 {
					passes += 1
				}
				position += 100
			}
			if position == 0 {
				passes += 1
			}

		}
		if position == 0 {
			stops += 1
		}
	}
	return solve(stops, passes)
}
