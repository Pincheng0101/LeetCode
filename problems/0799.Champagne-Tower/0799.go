package p0799

func champagneTower(poured int, query_row int, query_glass int) float64 {
	var tower [101][101]float64
	tower[0][0] = float64(poured)
	for row := 0; row <= query_row; row++ {
		for i := 0; i < row; i++ {
			if tower[row-1][i] <= 1 {
				continue
			}
			var halfOverflowWater float64 = (tower[row-1][i] - 1) / 2
			tower[row][i] += halfOverflowWater
			tower[row][i+1] += halfOverflowWater
		}
	}
	return min(1, tower[query_row][query_glass])
}

func min(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}
