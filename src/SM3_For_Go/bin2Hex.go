package SM3_For_Go

func bin2Hex(x []int) string {

	var result = ""
	var hexArr []int
	var k int

	for k = 0; k < (len(x) / 4); k++ {
		hexArr = append(hexArr, 0) //将16进制数组拓展到8位
	}

	for k = 0; k < len(hexArr); k++ {
		hexArr[k] = 8*x[4*k] + 4*x[(4*k)+1] + 2*x[(4*k)+2] + x[(4*k)+3] //2进制数组转16进制数组
	}
	for k = 0; k < len(hexArr); k++ { //十六进制数组转字符串
		switch hexArr[k] {
		case 0:
			result = result + "0"
		case 1:
			result = result + "1"
		case 2:
			result = result + "2"
		case 3:
			result = result + "3"
		case 4:
			result = result + "4"
		case 5:
			result = result + "5"
		case 6:
			result = result + "6"
		case 7:
			result = result + "7"
		case 8:
			result = result + "8"
		case 9:
			result = result + "9"
		case 10:
			result = result + "a"
		case 11:
			result = result + "b"
		case 12:
			result = result + "c"
		case 13:
			result = result + "d"
		case 14:
			result = result + "e"
		case 15:
			result = result + "f"
		}
	}
	return result
}
