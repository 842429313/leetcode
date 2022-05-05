package main

//超时间
// func recursive_nums(nums []int, k int, product int, num_temp *int) int {
// 	for i := 0; i < len(nums); i++ {
// 		temp := 0
// 		if product == 0 {
// 			temp = product + nums[i]
// 		} else {
// 			temp = product * nums[i]
// 		}
// 		if temp < k {
// 			*num_temp = *num_temp + 1
// 			if len(nums) != 1 {
// 				if product == 0 {
// 					recursive_nums(nums[i+1:], k, product+nums[i], num_temp)
// 				} else {
// 					recursive_nums(nums[i+1:], k, product*nums[i], num_temp)
// 				}

// 			}
// 			if product != 0 {
// 				break
// 			}
// 		} else {
// 			break
// 		}
// 	}
// 	return 0
// }

// func numSubarrayProductLessThanK(nums []int, k int) int {
// 	var num int = 0
// 	recursive_nums(nums, k, 0, &num)
// 	return num
// }

// func main() {
// 	nums := []int{10, 9, 10, 4, 3, 8, 3, 3, 6, 2, 10, 10, 9, 3}
// 	var k int = 19
// 	result := numSubarrayProductLessThanK(nums, k)
// 	fmt.Println((result))
// }
