package main

import (
	"time"
)

func main() {
	// switch x := math.Sqrt(4); x {
	// case 2:
	// 	fmt.Println("resultado é 2")
	// default:
	// 	fmt.Println("Deu erro")
	// }

}

func isWeekend(x time.Time) bool {
	switch x.Weekday() {
	case time.Sunday, time.Saturday:
		return true
	default:
		return false
	}
}

// func isWeekend(x time.Time) bool {
// 	switch {
// 	case x.Weekday() > 0 && x.Weekday() < 6:
// 		return false
// 	default:
// 		return true
// 	}

// switch {
// case x == 1:
// 	fmt.Println(1)
// 	fallthrough
// case "äbc" == "foo":
// 	fmt.Println(2)
// default:
// 	fmt.Println("Outra coisa")
// }

// }
