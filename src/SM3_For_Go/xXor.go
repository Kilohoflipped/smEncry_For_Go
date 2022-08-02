package SM3_For_Go

func xXor(x, y []int) []int { //异或
	var z []int
	var k int
	for k = 0; k < len(x); k++ {
		z = append(z, x[k]^y[k])
	}
	return z
}
