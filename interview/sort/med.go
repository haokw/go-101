package main

import "fmt"

func main() {
	nums1 := []int{1, 3}
	nums2 := []int{2}
	ans := findMedianSortedArrays(nums1, nums2)
	fmt.Println("nums1:", nums1)
	fmt.Println("nums2:", nums2)
	fmt.Println("ans:", ans)
}

func findMedianSortedArrays(nums1 []int, nums2 []int) (ans float64) {
	l1, l2 := len(nums1), len(nums2)

	// 中位位置 标识
	l := l1 + l2

	// 中位数组长度
	nums := make([]int, l)

	for i, i1, i2 := 0, 0, 0; i1 < l1 || i2 < l2; i++ {
		if i1 < l1 && i2 < l2 {
			// 每次只能添加一个
			if nums1[i1] < nums2[i2] {
				nums[i] = nums1[i1]
				i1++
			} else {
				nums[i] = nums2[i2]
				i2++
			}
		} else if i1 < l1 {
			nums[i] = nums1[i1]
			i1++
		} else if i2 < l2 {
			nums[i] = nums2[i2]
			i2++
		}
	}

	m := l / 2
	if l%2 == 0 {
		ans = float64(nums[m]+nums[m-1]) / 2
	} else {
		ans = float64(nums[m])
	}

	return
}
