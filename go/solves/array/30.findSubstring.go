package array

/*
 * @lc app=leetcode.cn id=30 lang=golang
 *
 * [30] 串联所有单词的子串
 */

// @lc code=start
func findSubstring(s string, words []string) []int {
	// 频率表
	wordCnt := make(map[string]int)
	for _, w := range words {
		wordCnt[w]++
	}
	wLen, totalLen, sLen := len(words[0]), len(words[0])*len(words), len(s)
	res := []int{}
	for offset := 0; offset < wLen && offset+totalLen <= sLen; offset++ { // 对不同的offset (0 ~ wLen-1), 注意右边界不超过s长度
		left, right := offset, offset+totalLen // 定义窗口 [left, right)
		cmpCnt := make(map[string]int)         // 用于比较的map
		for i := left; i+wLen <= right; i += wLen {
			cmpCnt[s[i:i+wLen]]++
		}
		if mapCmp(wordCnt, cmpCnt) {
			res = append(res, left)
		}
		left += wLen
		right += wLen
		for right <= sLen { // 循环从第二个开始
			leftStr := s[left-wLen : left]
			rightStr := s[right-wLen : right]
			cmpCnt[leftStr]--
			cmpCnt[rightStr]++
			if cmpCnt[leftStr] == 0 {
				delete(cmpCnt, leftStr)
			}
			if mapCmp(wordCnt, cmpCnt) {
				res = append(res, left)
			}
			left += wLen
			right += wLen
		}
	}
	return res
}
func mapCmp(m1, m2 map[string]int) bool {
	if len(m1) != len(m2) {
		return false
	}
	for k, v := range m1 {
		if m2[k] != v {
			return false
		}
	}
	return true
}

// @lc code=end
