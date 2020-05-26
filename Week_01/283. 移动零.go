package Week_01

func moveZeroes(nums []int) {

	ni := 0
	for _, num := range nums {
		if num != 0 {
			nums[ni] = num
			ni += 1
		}
	}
	for ni < len(nums) {
		nums[ni] = 0
		ni += 1
	}

}
