package SM3_For_Go

func sm3GGj(x, y, z []int, j int) []int { //布尔函数2
	var result []int
	if j >= 0 && j <= 15 {
		result = xXor(xXor(x, y), z)
	} else if j >= 16 && j <= 63 {
		result = xOr(xAnd(x, y), xAnd(xNot(x), z))
	}
	return result
}
