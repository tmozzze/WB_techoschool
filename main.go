package main

import (
	"fmt"

	stringunpacker "github.com/tmozzze/WB_techoschool/L2/L2_9/string_unpacker"
)

func main() {
	str := "qwe\\45"
	fmt.Println(stringunpacker.Unpack(str))
}
