package SM3_For_Go

func sm3FFj(x, y, z []int, j int) []int { //布尔函数1
	var result []int
	if j >= 0 && j <= 15 {
		result = xXor(xXor(x, y), z)
	} else if j >= 16 && j <= 63 {
		result = xOr(xOr(xAnd(x, y), xAnd(x, z)), xAnd(y, z))
	}
	return result
}
