package main

import "fmt"

func main() {
	v := 42 // change me!
	i := 42           // int
	f := 3.142        // float64
	g := 0.867 + 0.5i // complex128
	fmt.Printf("v is of Type %T\n", v)
	fmt.Printf("i is of Type %T\n", i)
	fmt.Printf("g is of Type %T\n", f)
	fmt.Printf("g is of Type %T\n", g)
}
