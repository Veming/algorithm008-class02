package Week_01

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	length := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[length] {
			length++
			nums[length] = nums[i]
		}
	}

	return length + 1
}
