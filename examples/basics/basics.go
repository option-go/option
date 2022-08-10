package main

import (
	"fmt"

	"github.com/option-go/option"
)

func main() {
	someInt := option.Some[int](3)
	noInt := option.None[int]()

	fmt.Printf("someInt.IsSome(): %t \n", someInt.IsSome())
	fmt.Printf("someInt.IsNone(): %t \n", someInt.IsNone())

	fmt.Printf("noInt.IsSome(): %t \n", noInt.IsSome())
	fmt.Printf("noInt.IsNone(): %t \n", noInt.IsNone())

	fmt.Printf("someInt.Unwrap(): %d \n", someInt.Unwrap())

	// this would panic!
	// fmt.Printf("noInt.Unwrap(): %d", noInt.Unwrap())
}
