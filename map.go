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
