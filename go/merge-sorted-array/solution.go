package solution

func merge(nums1 []int, m int, nums2 []int, n int) {
	last := m + n - 1
	n--
	m--

	for m >= 0 && n >= 0 {
		if nums1[m] > nums2[n] {
			nums1[last] = nums1[m]
			m--
		} else {
			nums1[last] = nums2[n]
			n--
		}
		last--
	}

	for n >= 0 {
		nums1[last] = nums2[n]
		n--
		last--
	}
}
