package main

// import (
// 	"fmt"
// 	"math/rand"
// 	"time"
// )

// type Solution map[int][]int

// func Constructor(nums []int) Solution {
// 	pos := map[int][]int{}
// 	for i, v := range nums {
// 		pos[v] = append(pos[v], i)
// 	}
// 	return pos
// }

// func (pos Solution) Pick(target int) int {
// 	indices := pos[target]
// 	return indices[rand.Intn(len(indices))]
// }

// func main() {
// 	rand.Seed(time.Now().UnixNano())
// 	nums := []int{1, 2, 3, 3, 3}
// 	obj := Constructor(nums)
// 	param_1 := obj.Pick(3)
// 	fmt.Println(param_1)
// }
