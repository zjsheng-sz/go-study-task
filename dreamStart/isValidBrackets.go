package dreamstart

func isValid(s string) bool {

	numMap := map[byte]byte{
		'}': '{',
		')': '(',
		']': '[',
	}

	stack := []byte{}

	for _, item := range s {

		if v, ok := numMap[byte(item)]; ok {

			if len(stack) > 0 && v == stack[len(stack)-1] {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}

		} else {

			stack = append(stack, byte(item))
		}

	}

	if len(stack) > 0 {
		return false
	}

	return true

}
