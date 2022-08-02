package SM3_For_Go

func sm3P0(x []int) []int { //置换函数1
	var result []int
	result = xXor(x, rotLeftBin(x, 9%32))
	result = xXor(result, rotLeftBin(x, 17%32))
	return result
}
