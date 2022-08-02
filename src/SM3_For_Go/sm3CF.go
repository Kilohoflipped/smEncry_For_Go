package SM3_For_Go

func sm3CF(V []string, B []int) stringsArr {
	var kj int
	var k int
	var V0 []string
	for k = 0; k < len(V); k++ { //将此时的字寄存器初始值记录下来
		V0 = append(V0, V[k]+"")
	}

	var temp []int            //为了防止语句过长使用中间变量
	var W [68]intsArr         //W的结构体数组，一个结构体里是长度为32的数组，68个拓展字1
	var W1 [64]intsArr        //W_1的结构体数组，一个结构体里是长度为32的数组，64个拓展字1
	for k = 0; k <= 15; k++ { //0-15个拓展字1
		W[k].data = B[(k * 32) : (k+1)*32]
	}

	for k = 16; k <= 67; k++ { //16-67个拓展字1
		temp = xXor(W[k-16].data, W[k-9].data)
		temp = xXor(temp, rotLeftBin(W[k-3].data, 15))
		temp = sm3P1(temp)
		temp = xXor(temp, rotLeftBin(W[k-13].data, 7))
		W[k].data = xXor(temp, W[k-6].data)
	}

	for k = 0; k <= 63; k++ { //0-63个拓展字2
		W1[k].data = xXor(W[k].data, W[k+4].data)
	}

	var VBinArr []intsArr
	for k = 0; k < len(V); k++ {
		VBinArr = append(VBinArr, intsArr{data: hex2Bin(V[k])}) //8个字寄存器的二进制表示
	}

	var SS1, SS2, TT1, TT2 []int
	var temptj []int

	for k = 0; k <= 63; k++ {
		if k <= 15 {
			temptj = []int{0, 1, 1, 1, 1, 0, 0, 1, 1, 1, 0, 0, 1, 1, 0, 0, 0, 1, 0, 0, 0, 1, 0, 1, 0, 0, 0, 1, 1, 0, 0, 1} //0x79cc4519
		} else {
			temptj = []int{0, 1, 1, 1, 1, 0, 1, 0, 1, 0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 0, 1, 0} //0x7a879d8a
		}
		temp = rotLeftBin(VBinArr[0].data, 12)
		temp = xPluxMod232(temp, VBinArr[4].data)
		temp = xPluxMod232(temp, rotLeftBin(temptj, k%32))
		SS1 = rotLeftBin(temp, 7%32) //A循环左移12位+E+Tj循环左移j位

		SS2 = xXor(SS1, rotLeftBin(VBinArr[0].data, 12)) //SS1与A循环左移12位求异或

		temp = xPluxMod232(sm3FFj(VBinArr[0].data, VBinArr[1].data, VBinArr[2].data, k), VBinArr[3].data)
		temp = xPluxMod232(temp, SS2)
		TT1 = xPluxMod232(temp, W1[k].data)

		temp = xPluxMod232(sm3GGj(VBinArr[4].data, VBinArr[5].data, VBinArr[6].data, k), VBinArr[7].data)
		temp = xPluxMod232(temp, SS1)
		TT2 = xPluxMod232(temp, W[k].data)

		V[3] = V[2]                                     //D <- C
		V[2] = bin2Hex(rotLeftBin(VBinArr[1].data, 9))  //C <- B<<<9
		V[1] = V[0]                                     //B <- A
		V[0] = bin2Hex(TT1)                             //A<-TT1
		V[7] = V[6]                                     //H <- G
		V[6] = bin2Hex(rotLeftBin(VBinArr[5].data, 19)) //G <- F<<<19
		V[5] = V[4]                                     //F <- E
		V[4] = bin2Hex(sm3P0(TT2))                      //E <- PP0(TT2)
		for kj = 0; kj < len(V); kj++ {
			VBinArr[kj] = intsArr{data: hex2Bin(V[kj])} //8个字寄存器的二进制表示
		}
	}
	var V1 stringsArr
	V1.data = []string{"", "", "", "", "", "", "", ""}
	for k = 0; k < len(V1.data); k++ {
		temp = xXor(hex2Bin(V[k]), hex2Bin(V0[k]))
		V1.data[k] = bin2Hex(temp)
	}
	return V1
}
