package gomath

func Sum(nums []int64) int64 {
	var res int64
	for _, v := range nums {
		res += res + v
	}

	return res
}