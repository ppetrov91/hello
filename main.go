package main

import (
	"fmt"

	"github.com/ppetrov91/hello/mascot"

	"rsc.io/quote"
)

func main() {
	fmt.Println("Hello, world.")
	fmt.Println(mascot.BestMascot())
	fmt.Println(quote.Go())
}
