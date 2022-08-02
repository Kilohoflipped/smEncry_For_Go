package SM3_For_Go

func sm3Hash(msg string) string {
	type stringsArr struct { //字符串数组结构体
		data []string
	}

	type intsArr struct { //整形数组结构体
		data []int
	}
	//字寄存器初始值
	var IV = []string{"7380166f", "4914b2b9", "172442d7", "da8a0600", "a96f30bc", "163138aa", "e38dee4d", "b0fb0e4e"}
	var msgBinArr = hex2Bin(msg)     //将十六进制字符串转为二进制数组存储
	var len0 = len(msgBinArr)        //初始消息的原比特长度
	var remainder = len0 % 512       //初始消息的原比特长度mod512的值(按512比特分组，最后一组比特数目)
	var k int                        //各种for循环的索引值
	msgBinArr = append(msgBinArr, 1) //在消息末尾追加一个比特1
	remainder += 1                   //消息末尾追加了1之后最后一组比特长度也+1
	var bitsLengthPretd = 512 - 64   //消息经过预处理之后，除开最后64位用于存储len1的二进制表示，剩下需要恰是这么多位(448)
	if remainder > bitsLengthPretd {
		bitsLengthPretd += 512 //如果消息末尾追加了1之后比特长度比448更大，即最后没有64比特的空间来存储len1的二进制表示值，则直接追加一组512个0
	}
	for k = remainder; k < bitsLengthPretd; k++ {
		msgBinArr = append(msgBinArr, 0) //追加0
	}

	var len0BinArr = dec2Bin(len0)                   //得到初始消息原比特长度值的二进制数组表示
	for k = len(len0BinArr) - 1; k < (64 - 1); k++ { //因为是下标所以全部减一
		len0BinArr = append([]int{0}, len0BinArr...) //把长度变成64
	}

	msgBinArr = append(msgBinArr, len0BinArr...) //预处理完毕的消息
	var Bits512Count = len(msgBinArr) / 512      //预处理完毕的消息按512bit分组的组数目

	//接下来把预处理完毕的消息按512bit进行分组在结构体数组里
	var B []intsArr //B的结构体数组，一个结构体里是长度为512的数组
	for k = 0; k < Bits512Count; k++ {
		B = append(B, intsArr{data: msgBinArr[(k * 512):((k + 1) * 512)]}) //把消息按512分组填充到B中
	}
	//fmt.Println(B)
	//fmt.Println(bin2Hex(B[0].data))
	var V []stringsArr //V的结构体数组，每个结构体里都是8个字寄存器字符串数组

	V = append(V, stringsArr{data: IV}) //初始化V
	for k = 0; k < Bits512Count; k++ {
		V = append(V, sm3CF(V[k].data, B[k].data)) //直接把字符串数组和二进制数传进去
	}
	var VE = V[k].data
	var result = ""
	for k = 0; k < len(VE); k++ {
		result = result + VE[k]
	}
	return result
}
