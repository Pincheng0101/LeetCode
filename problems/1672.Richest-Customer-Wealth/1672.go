package p1672

func maximumWealth(accounts [][]int) int {
	maximumWealth := 0
	for i := 0; i < len(accounts); i++ {
		sumWealth := 0
		for j := 0; j < len(accounts[i]); j++ {
			sumWealth += accounts[i][j]
		}
		if sumWealth > maximumWealth {
			maximumWealth = sumWealth
		}
	}
	return maximumWealth
}
