// +build OMIT

package main

func main() { // OMIT START
	switch x := 5; x {
	default:
		println(x)
		//fallthrough
	case 5:
		x += 10
		println(x)

		if x >= 15 {
			break
		}

		fallthrough // 必須是 case 區塊的最後一行語句
	case 6:
		x += 20
		println(x)
		//fallthrough 		// 如果在此繼續 fallthrough, 不會執行 default, 完全按源碼順序
		//   					將導致錯誤 "cannot fallthrough final case in switch"
	}
} // OMIT END
