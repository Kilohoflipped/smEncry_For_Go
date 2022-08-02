package SM3_For_Go

func xOr(x, y []int) []int { //或
	var z []int
	var k int
	for k = 0; k < len(x); k++ {
		z = append(z, x[k]|y[k])
	}
	return z
}
