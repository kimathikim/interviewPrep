package main

import "fmt"

func twoSum(nums []int, target int) []int {
    length := len(nums)
    for i := 0; i < length; i++{
        for j := i+1; j < length; j++{
            if nums[i] + nums[j] == target{return []int{i, j}}
        }      
    }
    return []int{0, 0}
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	result := twoSum(nums, target)
	fmt.Println(result) // Output: [0 1]

	nums = []int{3,2,4}
	target = 6
	result = twoSum(nums, target)
	fmt.Println(result) // Output: [1 2]

	nums = []int{3,3}
	target = 6
	result = twoSum(nums, target)
	fmt.Println(result) // Output: [0 1]

}

