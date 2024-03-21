package dynamic

func CutCable(price []int, leng int) int {
	matrix := make([]int, leng+1)
	for i := 0; i < leng+1; i++ {
		for j := 0; j < len(price); j++ {
			if j <= i {
				matrix[i] = max(matrix[i], matrix[i-j]+price[j])
			}
		}
	}
	return matrix[leng]
}
