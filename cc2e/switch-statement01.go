// +build OMIT

package main

func main() { // OMIT START
	a, b, c, x := 1, 2, 3, 2
	switch x {
	case a, b:
		println("a | b")
	case c:
		println("c")
	case 4:
		println("d")
	default:
		println("z")
	}

	switch {
	case x > 2:
		println("x>2")
	case x == 2:
		println("x=2")
	case x < 2:
		println("x<2")
	default:
		println("?")
	}
} // OMIT END
