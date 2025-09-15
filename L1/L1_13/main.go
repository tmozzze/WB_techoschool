package main

import "fmt"

func main() {
	a := 3
	b := 6

	num1, num2 := SwitchNums(a, b)

	fmt.Println(num1, num2)
}
func SwitchNums(a, b int) (int, int) {
	a = a ^ b
	b = a ^ b
	a = a ^ b

	return a, b

}
