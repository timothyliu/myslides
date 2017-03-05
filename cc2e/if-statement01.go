// +build OMIT
package main

import (
	"log"
	"strconv"
)

func main() { // OMIT START
	s := "x"
	if n, err := strconv.ParseInt(s, 10, 64); err != nil {
		log.Fatalln("invalid number: ", s, n)
	}
	println("ok")
} // OMIT END
