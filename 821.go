package main

import (
	"fmt"
	"math"
)

func shortestToChar(s string, c byte) []int {
	answer := make([]int, len(s))
	var last_direction int = 0
	for i := 0; i < len(answer); i++ {
		if i == 0 {
			answer[i] = len(s)
		} else if last_direction == 0 && answer[i-1] != 0 {
			answer[i] = answer[i-1] - 1
		} else {
			last_direction = 1
			answer[i] = answer[i-1] + 1
		}
		for j := i; j < len(s); j++ {
			if (s[j] == c) && math.Abs(float64(j-i)) < float64(answer[i]) {
				answer[i] = int(math.Abs(float64(j - i)))
				if (j - i) >= 0 {
					last_direction = 0
					break
				} else {
					last_direction = 1
				}
			}
		}
	}
	return answer
}

func main() {
	s := "loveleetcode"
	var c byte = 'e'
	fmt.Print(shortestToChar(s, c))

}
