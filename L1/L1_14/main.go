package main

import "fmt"

func main() {
	i := 10
	s := "sas"
	b := true
	ch := make(chan interface{})
	f := .1

	fmt.Println(CalcType(i))
	fmt.Println(CalcType(s))
	fmt.Println(CalcType(b))
	fmt.Println(CalcType(ch))
	fmt.Println(CalcType(f))
}

func CalcType(v interface{}) string {
	switch v.(type) {
	case int:
		return "integer"
	case string:
		return "string"
	case bool:
		return "boolean"
	case chan interface{}:
		return "channel of interface"
	}
	return "HZ cho tam"
}
