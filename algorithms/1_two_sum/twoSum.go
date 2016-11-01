package main

import "fmt"

func main() {
	nums := []int{2, 4, 6, 7}
	fmt.Println(twoSum(nums, 13))
	fmt.Println(twoSum(nums, 17))
}

func twoSum(nums []int, target int) []int {
    for i := 0; i < len(nums) - 1; i++ {
        for j := i + 1; j < len(nums); j++ {
            if nums[i] + nums[j] == target {
                return []int{i, j}
            }
        }
    }
    return nil
}
