package p0733

func floodFill(image [][]int, sr int, sc int, color int) [][]int {

	fill(image, sr, sc, color, image[sr][sc])

	return image
}

func fill(image [][]int, sr int, sc int, color int, oldColor int) {
	if sr < 0 || sc < 0 || sr >= len(image) || sc >= len(image[0]) {
		return
	}

	if image[sr][sc] != oldColor || image[sr][sc] == color {
		return
	}

	image[sr][sc] = color

	fill(image, sr+1, sc, color, oldColor)
	fill(image, sr-1, sc, color, oldColor)
	fill(image, sr, sc+1, color, oldColor)
	fill(image, sr, sc-1, color, oldColor)
}
