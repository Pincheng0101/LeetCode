package p0045

func jump_1(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	last := nums[0]
	i, n := 1, 1
	for last < len(nums)-1 {
		maxJump := 0
		for ; i < len(nums) && i < last+1; i++ {
			if i+nums[i] > maxJump {
				maxJump = i + nums[i]
			}
		}
		last = maxJump
		n++
	}
	return n
}

func jump_2(nums []int) int {
	n := 0
	farest, end := 0, 0
	for i := 0; i < len(nums)-1; i++ {
		if i+nums[i] > farest {
			farest = i + nums[i]
		}
		if i == end {
			end = farest
			n++
		}
	}
	return n
}
