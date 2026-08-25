package main

import "fmt"

func main() {
	var nums [4]int

	// length of array
	fmt.Println(len(nums))

	nums[0] = 1
	nums[1] = 2
	fmt.Println(nums)

	for i := 0; i < len(nums); i++ {
		fmt.Println(nums[i])
	}

	// to declare it in single line
	nums1 := [3]int{1, 2, 3}
	fmt.Println(nums1)

	// 2d arrays
	nums2 := [2][2]int{{2, 1}, {3, 4}}
	fmt.Println(nums2[0][1])
}

// - fixed siz , that is predictable
// - memory optimization
// - constant time access
