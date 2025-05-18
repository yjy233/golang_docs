package main

import (
	"fmt"
	_ "unsafe"

	_ "github.com/yjy233/golang_docs/cool_facts/golinkname/lib"
)

//go:linkname fA github.com/yjy233/golang_docs/cool_facts/golinkname/lib.funcA
func fA(a int)

func main() {

	fmt.Println("jyjy")
	fA(123)
}
