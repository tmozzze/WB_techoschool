package l122

import (
	"fmt"
	"log"
	"math/big"
)

func Calculate() {
	var aStr, bStr string
	fmt.Print("Enter A num: ")
	fmt.Scan(&aStr)
	fmt.Print("Enter B num: ")
	fmt.Scan(&bStr)

	a := new(big.Int)
	b := new(big.Int)

	_, ok := a.SetString(aStr, 10)
	if !ok {
		log.Printf("incorrect A num: %s\n", aStr)
		return
	}
	_, ok = b.SetString(bStr, 10)
	if !ok {
		log.Printf("incorrect B num: %s\n", bStr)
		return
	}

	choice := 0
	fmt.Print("What to do?\n 1. Sum\n 2. Diff\n 3. Multiply\n 4. Div\n 5. Exit\n")
	fmt.Print("Your choice: ")
	fmt.Scan(&choice)

	switch choice {
	case 1:
		sum := new(big.Int).Add(a, b)
		fmt.Println("Sum A and B equal:", sum)
	case 2:
		diff := new(big.Int).Sub(a, b)
		fmt.Println("Diff A and B equal:", diff)
	case 3:
		mult := new(big.Int).Mul(a, b)
		fmt.Println("Multiply A and B equal:", mult)
	case 4:
		if b.Cmp(big.NewInt(0)) != 0 {
			div := new(big.Int).Div(a, b)
			fmt.Println("Diff A and B equal:", div)
		} else {
			log.Println("cannot be divided by zero")
		}
	case 5:
		fmt.Println("Shutting down the program...")
		return
	default:
		fmt.Println("Incorrect choice. Restarting...")
		Calculate()
	}

}
