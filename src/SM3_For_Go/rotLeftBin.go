package SM3_For_Go

func rotLeftBin(x []int, n int) []int { //循环左移函数
	var k int
	var arrLefted []int
	var leftConst = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31}
	for k = 0; k < 32; k++ { //移位后数组对应数字的索引
		if leftConst[k]+(n%32) <= 31 {
			leftConst[k] += n % 32
		} else if leftConst[k]+(n%32) > 31 {
			leftConst[k] = leftConst[k] + (n % 32) - 32
		}
	}
	for k = 0; k <= 31; k++ {
		arrLefted = append(arrLefted, 0) //把新数组延长至32比特
	}
	for k = 0; k <= 31; k++ {
		arrLefted[k] = x[leftConst[k]] //构造移位后的数组
	}
	return arrLefted
}
