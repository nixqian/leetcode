package backtrack

import (
	"fmt"
	"strings"
)

/*
 * @lc app=leetcode.cn id=93 lang=golang
 *
 * [93] 复原 IP 地址
 */

// @lc code=start
func restoreIpAddresses(s string) []string {
	res := []string{}
	path := []string{}
	var bp func(dotCnt int, start int)
	bp = func(dotCnt, start int) {
		if start == len(s) {
			return
		}
		if dotCnt == 3 && ipGapIsValid(s[start:]) {
			fmt.Println(start)
			res = append(res, strings.Join(path, ".")+"."+s[start:])
			return
		}
		for i := start; i < len(s) && i < start+3; i++ {
			if ipGapIsValid(s[start : i+1]) {
				path = append(path, s[start:i+1])
				bp(dotCnt+1, i+1)
				path = path[:len(path)-1]
			}
		}
	}
	bp(0, 0)
	return res
}

func ipGapIsValid(s string) bool {
	if s[0] == '0' && len(s) > 1 {
		return false
	}
	n := 0
	for _, c := range s {
		n = n*10 + int(c-'0')
	}
	return n <= 255
}

// @lc code=end
