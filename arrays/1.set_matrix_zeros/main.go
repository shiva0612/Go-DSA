package main

func main() {
	// [[1,1,1],[1,0,1],[1,1,1]]
	ip := [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}

	// setZeroes1(ip)
	setZeroes2(ip)

	//checkout one last method where we use 1strowCOl of matrix instead of seperate rowCol arrays

}

// ----------------------------------------------------------------
func setZeroes1(matrix [][]int) {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 0 {
				matrix[i][j] = -1
			}
		}
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == -1 {
				setRowColtoZeros(matrix, i, j)
			}
		}
	}
	replacewithzero(matrix)

}

func setRowColtoZeros(matrix [][]int, r, c int) {
	for i := 0; i < len(matrix[r]); i++ {
		if matrix[r][i] == -1 {
			continue
		}
		matrix[r][i] = 0
	}

	for i := 0; i < len(matrix); i++ {
		if matrix[i][c] == -1 {
			continue
		}
		matrix[i][c] = 0
	}
}
func replacewithzero(matrix [][]int) {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == -1 {
				matrix[i][j] = 0
			}
		}
	}
}

// ----------------------------------------------------------------

func setZeroes2(matrix [][]int) {
	row := make([]int, len(matrix))
	col := make([]int, len(matrix[0]))

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				row[i] = -1
				col[j] = -1
			}
		}
	}

	for i := 0; i < len(row); i++ {
		if row[i] == -1 {
			setRowZeros(matrix, i)
		}
	}
	for i := 0; i < len(col); i++ {
		if col[i] == -1 {
			setColsZeros(matrix, i)
		}
	}
}

func setRowZeros(matrix [][]int, r int) {
	for i := 0; i < len(matrix[0]); i++ {
		matrix[r][i] = 0
	}
}
func setColsZeros(matrix [][]int, c int) {
	for i := 0; i < len(matrix); i++ {
		matrix[i][c] = 0
	}
}
