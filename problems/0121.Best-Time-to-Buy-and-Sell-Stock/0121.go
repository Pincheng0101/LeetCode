package p0121

func maxProfit(prices []int) int {
	maxProfit := 0
	buyPrice := prices[0]
	for i := 1; i < len(prices); i++ {
		if buyPrice > prices[i] {
			buyPrice = prices[i]
		} else if maxProfit < prices[i]-buyPrice {
			maxProfit = prices[i] - buyPrice
		}
	}
	return maxProfit
}
