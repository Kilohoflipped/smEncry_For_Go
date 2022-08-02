package SM3_For_Go

func xPluxMod232(x, y []int) []int { //模2^32比特加法运算
	var z []int
	var k int
	var flag = true

	for k = 0; k < len(x); k++ {
		z = append(z, x[k]+y[k]) //先以整型相加
	}

	for true {
		flag = true
		if z[0] == 2 {
			z[0] = 0 //最高位进位时取模
		}
		for k = 1; k < len(z); k++ {
			if z[k] == 2 {
				z[k] = 0
				z[k-1] += 1
				flag = false
			}
		}
		if flag {
			break //如果非最高位的数字没有2出现，说明已经迭代完成
		}
	}
	return z
}
