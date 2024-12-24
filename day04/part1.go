package day04

func partOne(content [][]rune) int {
	sum := 0
	for i := range content {
		for j := range content[i] {
			if checkTopLeft(content, i, j) {
				sum++
			}
			if checkTop(content, i, j) {
				sum++
			}
			if checkTopRight(content, i, j) {
				sum++
			}
			if checkLeft(content, i, j) {
				sum++
			}
		}
	}
	return sum
}

func checkTopLeft(content [][]rune, i, j int) bool {
	if i < 3 || j < 3 {
		return false
	}
	return isValid([]rune{content[i][j], content[i-1][j-1], content[i-2][j-2], content[i-3][j-3]})
}

func checkTop(content [][]rune, i, j int) bool {
	if i < 3 {
		return false
	}
	return isValid([]rune{content[i][j], content[i-1][j], content[i-2][j], content[i-3][j]})
}

func checkTopRight(content [][]rune, i, j int) bool {
	if i < 3 || j >= len(content[i])-3 {
		return false
	}
	return isValid([]rune{content[i][j], content[i-1][j+1], content[i-2][j+2], content[i-3][j+3]})
}

func checkLeft(content [][]rune, i, j int) bool {
	if j < 3 {
		return false
	}
	return isValid([]rune{content[i][j], content[i][j-1], content[i][j-2], content[i][j-3]})
}

func isValid(word []rune) bool {
	return isXMAS(word) || isSAMX(word)
}

func isXMAS(word []rune) bool {
	return len(word) == 4 && word[0] == 'X' && word[1] == 'M' && word[2] == 'A' && word[3] == 'S'
}
func isSAMX(word []rune) bool {
	return len(word) == 4 && (word[0] == 'S' && word[1] == 'A' && word[2] == 'M' && word[3] == 'X')
}
