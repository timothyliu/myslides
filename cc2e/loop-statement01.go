// +build OMIT

package main

func main() { // OMIT START
	for i := 0; i < 3; i++ {
		println(i)
	}

	x := 8
	for x < 10 { // 類似  while x < 10 {} 或  for ; x< 10 ; {}
		x++
		println(x)
	}

	data := [3]string{"a", "b", "c"}
	for i, s := range data { // 類似 foreach
		println(i, s)
	}

	for { // 類似  while true {} 或  for true {}
		break
	}
} // OMIT END
