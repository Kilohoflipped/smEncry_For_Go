package SM3_For_Go

func hex2Bin(x string) []int {
	var binArr []int
	var k int
	for k = 0; k < len(x); k++ {
		switch x[k] {
		case 48:
			binArr = append(binArr, []int{0, 0, 0, 0}...) //0
		case 49:
			binArr = append(binArr, []int{0, 0, 0, 1}...) //1
		case 50:
			binArr = append(binArr, []int{0, 0, 1, 0}...) //2
		case 51:
			binArr = append(binArr, []int{0, 0, 1, 1}...) //3
		case 52:
			binArr = append(binArr, []int{0, 1, 0, 0}...) //4
		case 53:
			binArr = append(binArr, []int{0, 1, 0, 1}...) //5
		case 54:
			binArr = append(binArr, []int{0, 1, 1, 0}...) //6
		case 55:
			binArr = append(binArr, []int{0, 1, 1, 1}...) //7
		case 56:
			binArr = append(binArr, []int{1, 0, 0, 0}...) //8
		case 57:
			binArr = append(binArr, []int{1, 0, 0, 1}...) //9
		case 65, 97:
			binArr = append(binArr, []int{1, 0, 1, 0}...) //a/A
		case 66, 98:
			binArr = append(binArr, []int{1, 0, 1, 1}...) //b/B
		case 67, 99:
			binArr = append(binArr, []int{1, 1, 0, 0}...) //c/C
		case 68, 100:
			binArr = append(binArr, []int{1, 1, 0, 1}...) //d/D
		case 69, 101:
			binArr = append(binArr, []int{1, 1, 1, 0}...) //e/E
		case 70, 102:
			binArr = append(binArr, []int{1, 1, 1, 1}...) //f/F
		}
	}
	return binArr
}
