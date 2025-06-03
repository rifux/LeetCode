package solution

func removeElement(nums []int, val int) (counter int) {
	for _, num := range nums {
		if num != val {
			nums[counter] = num
			counter++
		}
	}
	return counter
}
