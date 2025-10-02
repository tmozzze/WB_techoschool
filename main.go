package main

import "fmt"

func main() {
	var s = []string{}
	s = append(s, "1")
	s = append(s, "2")
	s = append(s, "3")
	modifySlice(s)
	fmt.Println(s)
}

func modifySlice(i []string) {
	i[0] = "3"
	i = append(i, "4")
	i[1] = "5"
	i = append(i, "6")
}
