package SM3_For_Go

func sm3P1(x []int) []int { //置换函数2
	var result []int
	result = xXor(x, rotLeftBin(x, 15%32))
	result = xXor(result, rotLeftBin(x, 23%32))
	return result
}
