package leetcode

import "fmt"

func V1canBeValid(s string, locked string) bool {

	n := len(s)

	if n%2 != 0 {
		return false
	}

	for i, j := 0, 1; j < n; i, j = i+2, j+2 {
		if (s[i] == ')' && locked[i] == '1') ||
			(s[j] == '(' && locked[j] == '1') ||
			(s[i] == '(' && locked[i] == '1' && s[j] == '(' && locked[j] == '1') ||
			(s[i] == ')' && locked[i] == '1' && s[j] == ')' && locked[j] == '1') {
			return false
		}
	}

	return true
}

func V2canBeValid(s string, locked string) bool {

	n := len(s)

	if n%2 != 0 {
		return false
	}

	opens := []int{}

	unlocked := []int{}

	for i := 0; i < n; i++ {

		if locked[i] == '0' {
			unlocked = append(unlocked, i)
		} else if s[i] == '(' {
			opens = append(opens, i)
		} else {
			
			if len(opens) > 0 {
				opens = opens[:len(opens)-1]
			} else if len(unlocked) > 0 {
				unlocked = unlocked[:len(unlocked)-1]
			} else {
				return false
			}
		}
	}

	for len(opens) > 0 && len(unlocked) > 0 {
		if unlocked[len(unlocked)-1] < opens[len(opens)-1] {
			return false
		}

		unlocked = unlocked[:len(unlocked)-1]
		opens = opens[:len(opens)-1]
	}

	return len(opens) == 0 && len(unlocked)%2 == 0
}

const (
	zero  byte = '0'
	one   byte = '1'
	left  byte = '('
	right byte = ')'
)

func canBeValid(s string, locked string) bool {

	count := 0
	open := 0

	for i := range len(s) {

		lock := locked[i]
		bracket := s[i]

		if lock == zero {
			
			if 2*count < open {
				count++
			}
			
			open++
			continue
		}

		if bracket == left {
			
			open++
			continue
		}

		if bracket == right {
			
			if open > 0 {
				
				open--
				
				if 2*count > open {
					count--
				}
				
				continue
			}
			
			return false
		}
	}

	return !(open > 0 && 2*count < open)
}

func CanBeValid() {

	s := "(()())"
	locked := "111111"

	fmt.Println(canBeValid(s, locked))
}
