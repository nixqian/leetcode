package strings

import "fmt"

/*
 * @lc app=leetcode.cn id=299 lang=golang
 *
 * [299] 猜数字游戏
 */

// @lc code=start
func getHint(secret string, guess string) string {
	aCnt := 0
	bCntList := [10][2]int{}
	for i := 0; i < len(secret); i++ {
		if secret[i] == guess[i] {
			aCnt++
		} else {
			bCntList[int(secret[i]-'0')][0]++
			bCntList[int(guess[i]-'0')][1]++
		}
	}
	bCnt := 0
	for _, ele := range bCntList {
		if ele[0] > ele[1] {
			bCnt += ele[1]
		} else {
			bCnt += ele[0]
		}
	}
	return fmt.Sprintf("%dA%dB", aCnt, bCnt)
}

// @lc code=end
