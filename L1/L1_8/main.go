package main

import "fmt"

func main() {
	var num int64
	num = 5

	fmt.Printf("num: %d, (%04b)\n", num, num)

	changeNum := ChangeBit(num, 1) // меняем на первый бит
	fmt.Printf("change 1 bit num: %d, (%04b)\n", changeNum, changeNum)
	changeNum = ChangeBit(changeNum, 2) // меняем второй бит
	fmt.Printf("change 2 bit num: %d, (%04b)\n", changeNum, changeNum)

	setNum := SetBit(num, 1, false) //обнуляем первый бит числа
	fmt.Printf("set 1 bit null num: %d, (%04b)\n", setNum, setNum)
}

func ChangeBit(num int64, bit uint) int64 {
	mask := int64(1) << (bit - 1)

	return num ^ mask // ксорим маску с числом
}

func SetBit(num int64, bit uint, flag bool) int64 {
	mask := int64(1) << (bit - 1)

	if flag {
		return num | mask
	}
	return num &^ mask
}
