package main

import "fmt"

func maxRotateFunction(nums []int) int {
	var result int = 0
	first_on := 0
	for i := 0; i < len(nums); i++ {
		temp := 0
		for j := 0; j < len(nums); j++ {
			temp = temp + (j+i)%len(nums)*nums[j]
		}
		fmt.Println(temp)

		if first_on == 0 {
			result = temp
			first_on = 1
		} else if temp > result {
			result = temp
		}
	}
	fmt.Println("\n")
	return result
}
func maxRotateFunction_2(nums []int) int {
	var result int = 0
	sum_nums := 0

	for i := 0; i < len(nums); i++ {
		sum_nums = sum_nums + nums[i]
	}
	for i := 0; i < len(nums); i++ {
		result = result + i*nums[i]
	}
	temp := result
	for i := 1; i < len(nums); i++ {
		if result < temp+sum_nums-len(nums)*nums[len(nums)-i] {
			result = temp + sum_nums - len(nums)*nums[len(nums)-i]

		}
		temp = temp + sum_nums - len(nums)*nums[len(nums)-i]
		fmt.Println((temp))

	}
	return result
}
func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, -10}
	fmt.Println(maxRotateFunction(nums) == maxRotateFunction_2(nums))
}
