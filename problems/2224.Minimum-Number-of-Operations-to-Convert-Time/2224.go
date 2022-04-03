package p2224

func convertTime(current string, correct string) int {
	currentSec := int(current[0]-0x30)*600 + int(current[1]-0x30)*60 + int(current[3]-0x30)*10 + int(current[4]-0x30)
	correctSec := int(correct[0]-0x30)*600 + int(correct[1]-0x30)*60 + int(correct[3]-0x30)*10 + int(correct[4]-0x30)

	diff := currentSec - correctSec
	ans := 0
	ans += diff / 60
	diff %= 60
	ans += diff / 15
	diff %= 15
	ans += diff / 5
	diff %= 5
	ans += diff

	if ans > 0 {
		return ans
	}
	return -ans
}
