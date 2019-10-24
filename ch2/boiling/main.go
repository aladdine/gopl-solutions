// boiling prints the boiling temperatue for water.
package main

import (
	"fmt"
)


const boilingF = 212.0 // package level declaration.

func main() {
	var f = boilingF
	var c = (f - 32) * 5 / 9
	fmt.Printf("boiling point = %gF or %g C\n", f, c)
}