package main

func main(){

}
func floodFill(image [][]int, startRow int, startColumn int, newColor int) [][]int {
	initialColor := image[startRow][startColumn]
	if initialColor != newColor {
		dfs(
			image,
			startRow,
			startColumn,
			initialColor,
			newColor,
		)
	}
	return image
}

func dfs(
	image [][]int,
	currentRow int,
	currentColumn int,
	initialColor int,
	newColor int,
) {
	if image[currentRow][currentColumn] != initialColor {
		return
	}
	image[currentRow][currentColumn] = newColor
	if currentRow >= 1 {
		dfs(
			image,
			currentRow-1,
			currentColumn,
			initialColor,
			newColor,
		)
	}
	if currentColumn >= 1 {
		dfs(
			image,
			currentRow,
			currentColumn-1,
			initialColor,
			newColor,
		)
	}
	if currentRow+1 < len(image) {
		dfs(
			image,
			currentRow+1,
			currentColumn,
			initialColor,
			newColor,
		)
	}
	if currentColumn+1 < len(image[0]) {
		dfs(
			image,
			currentRow,
			currentColumn+1,
			initialColor,
			newColor,
		)
	}
}

