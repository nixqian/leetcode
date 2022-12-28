package backtrack

func solveNQueens(n int) [][]string {
	res := [][]string{}
	path := [][]byte{}
	for i := 0; i < n; i++ {
		line := ""
		for j := 0; j < n; j++ {
			line = line + "."
		}
		path = append(path, []byte(line))
	}
	var backtrack func(n, level int)
	backtrack = func(n, level int) {
		if level == n {
			oneRes := []string{}
			for _, line := range path {
				oneRes = append(oneRes, string(line))
			}
			res = append(res, oneRes)
			return
		}
		for k := 0; k < n; k++ {

			if !isValid(level, k, n, path) {
				continue
			}
			path[level][k] = 'Q'
			backtrack(n, level+1)
			path[level][k] = '.'
		}
	}
	backtrack(n, 0)
	return res
}

func isValid(x, y, n int, board [][]byte) bool {
	// 横
	for j := y; j >= 0; j-- {
		if board[x][j] == 'Q' {
			return false
		}
	}
	// 竖
	for i := x; i >= 0; i-- {
		if board[i][y] == 'Q' {
			return false
		}
	}
	// 45度
	i, j := x, y
	for i >= 0 && j < n {
		if board[i][j] == 'Q' {
			return false
		}
		i--
		j++
	}
	// 135度
	i, j = x, y
	for i >= 0 && j >= 0 {
		if board[i][j] == 'Q' {
			return false
		}
		i--
		j--
	}
	return true
}
