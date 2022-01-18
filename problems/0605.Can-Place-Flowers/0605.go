package p0605

func canPlaceFlowers(flowerbed []int, n int) bool {
	for i := 0; i < len(flowerbed) && n >= 0; {
		if flowerbed[i] == 1 {
			i += 2
			continue
		}
		if i+1 == len(flowerbed) {
			n--
		} else if flowerbed[i+1] == 0 {
			n--
			i += 2
			continue
		}
		i++
	}

	return n <= 0
}
