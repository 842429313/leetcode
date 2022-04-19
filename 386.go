package main

// 方法是对的，但是超时

// func lexicalOrder(n int) []int {
// 	result := make([]int, n)
// 	for i := 1; i <= n; i++ {
// 		for j := 0; j < len(result); j++ {
// 			if result[j] == 0 {
// 				result[j] = i
// 				break
// 			}
// 			if compare(result[j], i) {
// 				a := make([]int, n)
// 				copy(a, result[j:])
// 				b := append(result[:j], i)
// 				result = append(b, a...)
// 				break
// 			}
// 		}

// 	}
// 	return result[:n]
// }
// func compare(a int, b int) bool {
// 	A := strconv.Itoa(a)
// 	B := strconv.Itoa(b)
// 	for i := 0; i < len(A) && i < len(B); i++ {
// 		if A[i] > B[i] {
// 			return true
// 		}
// 		if A[i] < B[i] {
// 			return false
// 		}

// 	}
// 	return false
// }

// func main() {
// 	fmt.Print(lexicalOrder(13))

// }
