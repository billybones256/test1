package main

import "fmt"

// Checker ...
func Checker(message string) int {
	maxLen := 0
	curLen := 0
	for _, ch := range message {
		if ch == 'o' {
			curLen++
		} else {
			if curLen > maxLen {
				maxLen = curLen
			}
			curLen = 0
		}
	}
	if curLen > maxLen {
		return curLen
	}
	return maxLen
}

func main() {
	fmt.Println(Checker("opopopopoooooop"))
}
