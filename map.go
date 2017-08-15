package main

// GenerateMap is function which we can use to generate map
func GenerateMap(x, y int) [][]int {
	gameMap := make([][]int, y)

	for i := 0; i < y; i++ {
		gameMap[i] = make([]int, x)
	}

	for i := 0; i < y; i++ {
		for j := 0; j < x; j++ {
			if i == 0 || i == y-1 || j == 0 || j == x-1 || (j%2 == 0 && i%2 == 0) {
				gameMap[i][j] = 1
			} else {
				gameMap[i][j] = 0
			}
		}
	}

	return gameMap
}

type point struct {
	X int
	Y int
}

func calculateStartPoint(x, y, width, height, sliceWidth, sliceHeight int) point {
	rx := sliceWidth / 2
	ry := sliceHeight / 2

	startPoint := point{}

	if x < rx {
		startPoint.X = 0
	} else if x+rx > width {
		startPoint.X = width - rx
	} else {
		startPoint.X = x - rx
	}

	if y < ry {
		startPoint.Y = 0
	} else if y+ry > height {
		startPoint.Y = height - ry
	} else {
		startPoint.Y = y - ry
	}

	return startPoint
}

func getPartOfArray(array [][]int, x, y, width, height, sliceWidth, sliceHeight int) [][]int {
	arrayPart := make([][]int, sliceHeight)

	for i := range arrayPart {
		arrayPart[i] = make([]int, sliceWidth)
	}

	startPoint := calculateStartPoint(x, y, width, height, sliceWidth, sliceHeight)
	tempX := startPoint.X
	tempY := startPoint.Y

	for i, row := range arrayPart {
		for j := range row {
			arrayPart[i][j] = array[tempY][tempX]
			tempX++
		}
		tempX = startPoint.X
		tempY++
	}

	return arrayPart
}
