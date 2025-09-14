package dreamstart

/*
输入：nums = [4,1,2,1,2]

输出：4

*/

func singleNumber(nums []int) int {

	// 
	m := map[int]bool{}

	for i := 0; i < len(nums); i++ {
		
		if m[nums[i]] {
			
			delete(m, nums[i])

		} else {

			m[nums[i]] = true
		}

	}
	
	for key , _ := range m {
		return  key
	}

	return 0
}
