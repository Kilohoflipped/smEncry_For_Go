package SM3_For_Go

func xNot(x []int) []int { //非
	var z []int
	var k int
	for k = 0; k < len(x); k++ {
		if x[k] == 1 {
			z = append(z, 0)
		} else {
			z = append(z, 1)
		}
	}
	return z
}
