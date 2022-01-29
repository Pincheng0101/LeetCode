package p0605

func canPlaceFlowers(flowerbed []int, n int) bool {
	size := len(flowerbed)
	for i := 0; i < size && n >= 0; i++ {
		if flowerbed[i] == 0 &&
			(i == 0 || flowerbed[i-1] == 0) &&
			(i == size-1 || flowerbed[i+1] == 0) {
			n--
			i++
		}
	}
	return n <= 0
}
