package main

import "fmt"

func numSubarrayProductLessThanK(nums []int, k int) int {
	var num int = 0
	i := 0
	product := 0
	j := 0
	for i != len(nums) {
		product = 0
		for j = i ; j < len(nums); j++ {
			if product == 0 {
				product = product + nums[j]
			} else {
				product = product * nums[j]
			}
			if product < k {
				num = num + 1
			} else {
				break
			}
		}
		j = 0
		i = i + 1
	}
	return num
}

func main() {
	nums := []int{10, 9, 10, 4, 3, 8, 3, 3, 6, 2, 10, 10, 9, 3}
	var k int = 19
	result := numSubarrayProductLessThanK(nums, k)
	fmt.Println((result))
}
