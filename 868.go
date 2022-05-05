package main

// import "fmt"

// func binaryGap(n int) int {
// 	distance := 0
// 	is_first := 0
// 	first_location := 0

// 	for i := 0; n != 0; i++ {
// 		if n%2 == 1 {
// 			if is_first == 0 {
// 				first_location = i
// 				is_first = 1
// 			} else if distance < i-first_location {
// 				distance = i - first_location
// 				first_location = i
// 			} else {
// 				first_location = i
// 			}

// 		}
// 		n = int(n / 2)
// 	}
// 	return distance
// }

// func main() {
// 	fmt.Println(binaryGap(15))
// }
