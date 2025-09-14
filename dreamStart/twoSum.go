package dreamstart

func twoSum1(nums []int, target int) []int {

	length := len(nums)

	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}

	return []int{}
}

func twoSum2(nums []int, target int) []int {

	numMap := map[int]int{}

	for i, v := range nums {

		
		if idx, ok := numMap[target-v]; ok {
			
			return  []int{idx, i}

		} else {

			numMap[v] = i
		}

	}
	return  []int{}

}
