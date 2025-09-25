package utils

func checkColumn(column int, grid [6][7]int) bool {

	//6 lignes 7 column

	for _, value := range grid[column] {
		if value == 0 {
			return false
		}
	}

	return true

}

func checkRow() {

}
