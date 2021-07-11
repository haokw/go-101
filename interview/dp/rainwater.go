func trap(height []int) int {
	len := len(height)

	left_max_arr := []int{}
	left_max := 0
	for i := 0; i < height; i++ {
		if height[i] > left_max {
			left_max_arr[i] = height[i]
		} else {
			left_max_arr[i] = left_max
		}
	}

	right_max_arr := []int{}
	right_max := 0
	for i := height - 1; i >= 0; i-- {
		if height[i] > right_max {
			right_max_arr[i] = height[i]
		} else {
			right_max_arr[i] = right_max
		}
	}

	for i, v := range height {
		ans += min(left_max_arr[i], right_max_arr[j]) - v
	}

	return
}
