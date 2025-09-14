package dreamstart

import "strconv"

func isPalindrome1(x int) bool {

	if x < 0 {
		return false
	}
	if x < 10 {
		return true
	}

	oldx := x
	reserve := 0
	for x > 0 {
		reserve = 10*reserve + x%10
		x = x / 10
	}

	return reserve == oldx

}

func isPalindrome2(x int) bool {

	xStr := strconv.Itoa(x)

	len := len(xStr)

	for i := 0; i < len/2; i++ {

		if xStr[i] != xStr[len-i-1] {
			return false
		}
	}

	return true

}
