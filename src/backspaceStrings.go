package main

// https://leetcode.com/explore/challenge/card/30-day-leetcoding-challenge/529/week-2/3291/

import "fmt"

func main() {
	matched := backspaceCompare(" ##a", "b")

	fmt.Println(matched)
}

func backspaceCompare(S string, T string) bool {
	buf1 := make([]byte, 0)
	for i := range S {
		if S[i] == '#' {
			if len(buf1) > 0 {
				buf1 = buf1[:len(buf1)-1]
			}
			continue
		}
		buf1 = append(buf1, S[i])
	}

	buf2 := make([]byte, 0)
	for i := range T {
		if T[i] == '#' {
			if len(buf2) > 0 {
				buf2 = buf2[:len(buf2)-1]
			}
			continue
		}
		buf2 = append(buf2, T[i])
	}

	return string(buf1) == string(buf2)
}
