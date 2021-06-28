package main

import (
	"fmt"
)

func t_twoSum() {
	a := twoSum([]int{-1, -2, -3, -4, -5}, -8)
	fmt.Println(a)
}

// func twoSum(nums []int, target int) []int {
// 	for k, v := range nums {
// 		for k2, v2 := range nums {
// 			if k == k2 {
// 				continue
// 			}
// 			if v+v2 == target {
// 				return []int{k, k2}
// 			}
// 		}
// 	}
// 	return []int{}
// }

// version 2
func twoSum(nums []int, target int) []int {
	// number : index
	seen_nums := map[int]int{}
	for k, v := range nums {
		remainder := target - v
		if index, ok := seen_nums[remainder]; ok {
			return []int{k, index}
		}
		seen_nums[v] = k
	}
	return []int{}
}
