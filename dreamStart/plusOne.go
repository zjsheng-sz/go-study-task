package dreamstart

func PlusOne(digits []int) []int {

	resNums := []int{}

	n := len(digits)

	j := 0

	digits[n-1] = digits[n-1] + 1

	for i := n - 1; i >= 0; i-- {

		resNums = append([]int{(digits[i] + j) % 10}, resNums...)

		if digits[i]+j >= 10 {
			j = 1
		} else {
			j = 0
		}

	}

	if j == 1 {

		resNums = append([]int{(1) % 10}, resNums...)

	}

	return resNums
}
