package solution1

func addBinary(a, b string) string {
	i, j := len(a)-1, len(b)-1
	var carry int
	var result []byte

	for i >= 0 || j >= 0 || carry > 0 {
		total := carry
		if i >= 0 {
			total += int(a[i] - '0')
			i--
		}
		if j >= 0 {
			total += int(b[j] - '0')
			j--
		}

		carry, result = total/2, append(result, byte(total%2+'0'))
	}

	for l, r := 0, len(result)-1; l < r; l, r = l+1, r-1 {
		result[l], result[r] = result[r], result[l]
	}

	return string(result)
}
