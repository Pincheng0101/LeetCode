package p0682

import "strconv"

func calPoints(ops []string) int {
	nums := []int{}
	for _, op := range ops {
		switch op {
		case "C":
			nums = nums[:len(nums)-1]
		case "D":
			newNum := 2 * nums[len(nums)-1]
			nums = append(nums, newNum)
		case "+":
			newNum := nums[len(nums)-1] + nums[len(nums)-2]
			nums = append(nums, newNum)
		default:
			v, _ := strconv.Atoi(op)
			nums = append(nums, v)
		}
	}
	ans := 0
	for _, v := range nums {
		ans += v
	}
	return ans
}
