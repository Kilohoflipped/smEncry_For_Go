package SM3_For_Go

func dec2Bin(x int) []int {
	var k, temp int
	var binArr, tempArr []int
	for k = x; k > 0; k /= 2 {
		temp = k % 2
		tempArr = append(tempArr, temp)
	}
	for k = 0; k < len(tempArr); k++ { //翻转数组
		binArr = append(binArr, tempArr[len(tempArr)-1-k])
	}
	return binArr
}
