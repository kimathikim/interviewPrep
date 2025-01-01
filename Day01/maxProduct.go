package main

import (
	"fmt"
	"sort"
)

func maximumProduct(nums []int) int {
	// get the length of the slice
	length := len(nums)
	// sort the numbers
	// i will use bubble sort for this
	/* for i := 0; i < length; i++{
	       for j := i+1; j < length; j++{
	           if nums[i] > nums[j]{

	               nums[i], nums[j] = nums[j], nums[i]
	           }
	       }
	   }
	*/
	// bubble sort was slow
	sort.Ints(nums)

	num := nums
	// check 3 cases. if the first and last are all positive, return the sum of the last three numbe
	max := num[length-1] * num[length-2] * num[length-3]
	fmt.Println(num)
	if num[0] >= 0 && num[length-1] > 0 {
		return num[length-1] * num[length-2] * num[length-3]
	}
	// if the first is negative and last positive return the the sum of the first 2 and the last value
	if num[0] < 0 && num[1] < 0 && num[length-1] > 0 {
		negUsed := num[0] * num[1] * num[length-1]
		if negUsed > max {
			return negUsed
		}
		return max
	}

	// if the first and the last are negative, return the sum of the first 3

	if num[0] < 0 && num[length-1] < 0 {
		return num[length-1] * num[length-2] * num[length-3]
	}
	return max
}

// white the main function taht will test the function
func main() {
	nums := []int{1, 2, 3, 4}
	fmt.Println(maximumProduct(nums))
}
