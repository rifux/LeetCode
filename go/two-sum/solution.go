package solution

func twoSum(nums []int, target int) []int {
	comparable := make(map[int]int)
	for i, num := range nums {
		j, ok := comparable[target-num]
		if ok {
			return []int{j, i}
		}
		comparable[num] = i
	}
	return []int{}
}
