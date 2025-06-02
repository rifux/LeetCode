package solution

func removeDuplicates(nums []int) (result int) {
	for im := 0; im < len(nums)-1; im++ {
		if nums[im] != nums[im+1] {
			nums[result+1] = nums[im+1]
			result++
		}
	}
	return result + 1
}
