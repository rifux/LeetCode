package solution2

import "strconv"

func addBinary(a string, b string) (c string) {
	var carry int

	for i := 0; i < max(len(a), len(b)); i++ {
		var curA, curB, curC int
		if i < len(a) {
			curA, _ = strconv.Atoi(string(a[len(a)-1-i]))
		}
		if i < len(b) {
			curB, _ = strconv.Atoi(string(b[len(b)-1-i]))
		}

		total := curA + curB + carry
		carry, curC = total/2, total%2

		c = strconv.Itoa(curC) + c
		/* ^^^
		This line is leading the whole approach to O(n2) time complexity,
		because every iteration we create new string of n length.

		If you want, you can use list of bytes that would be backward, so
		you will need to correct the order of bytes by loop.
		(../solution-1)
		*/
	}

	if carry == 1 {
		c = "1" + c
	}

	return
}
