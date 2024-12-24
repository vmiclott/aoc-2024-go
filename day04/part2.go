package day04

func partTwo(content [][]rune) int {
	sum := 0
	for i := 1; i < len(content)-1; i++ {
		for j := 1; j < len(content[i])-1; j++ {
			if content[i][j] == 'A' {
				if checkX(content, i, j) {
					sum++
				}
			}
		}
	}
	return sum
}

func checkX(content [][]rune, i int, j int) bool {
	return (content[i-1][j-1] == 'S' && content[i-1][j+1] == 'S' && content[i+1][j-1] == 'M' && content[i+1][j+1] == 'M') ||
		(content[i-1][j-1] == 'S' && content[i-1][j+1] == 'M' && content[i+1][j-1] == 'S' && content[i+1][j+1] == 'M') ||
		(content[i-1][j-1] == 'M' && content[i-1][j+1] == 'S' && content[i+1][j-1] == 'M' && content[i+1][j+1] == 'S') ||
		(content[i-1][j-1] == 'M' && content[i-1][j+1] == 'M' && content[i+1][j-1] == 'S' && content[i+1][j+1] == 'S')
}
