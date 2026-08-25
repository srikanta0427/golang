package main

import (
	"fmt"
	"slices"
)

func main() {
	var num []int
	if num == nil {
		fmt.Println(num)
	}

	var nums = make([]int, 5, 5)
	// capacity -> maximum numbers of elements can fit
	fmt.Println(cap(nums))

	nums = append(nums, 1)
	nums = append(nums, 2)
	nums = append(nums, 1)
	nums = append(nums, 2)
	nums = append(nums, 1)
	nums = append(nums, 2)
	nums[0] = 1
	nums[1] = 2
	nums[2] = 12
	fmt.Println(nums)
	fmt.Println(cap(nums))

	// also
	arr := []int{}
	arr = append(arr, 1)
	fmt.Println(cap(arr))
	fmt.Println(arr)

	var arr1 = make([]int, 0, 5)
	var arr3 = append(arr1, 2)
	arr1 = append(arr3, 4)
	var arr2 = make([]int, len(arr1))
	fmt.Println(arr3, arr1)

	// copy function
	copy(arr2, arr1)
	fmt.Println(arr2)

	fmt.Println(slices.Equal(arr1, arr2))

	// 2d arrays
	var arr4 = [][]int{{3, 4, 5}, {1, 2}}
	fmt.Println(arr4)
	for i := 0; i < len(arr4); i++ {
		for j := 0; j < len(arr4[i]); j++ {
			fmt.Print(arr4[i][j], " ")
		}
		fmt.Println()
	}
}
