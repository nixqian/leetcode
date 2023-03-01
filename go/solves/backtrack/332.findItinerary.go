package backtrack

import "sort"

/*
 * @lc app=leetcode.cn id=332 lang=golang
 *
 * [332] 重新安排行程
 */

// @lc code=start
func findItinerary(tickets [][]string) []string {
	ticketMap := make(map[string]map[string]int)
	for _, ticket := range tickets {
		_, ok := ticketMap[ticket[0]]
		if !ok {
			ticketMap[ticket[0]] = map[string]int{}
		}
		ticketMap[ticket[0]][ticket[1]]++
	}
	res := []string{"JFK"}
	var dfs func() bool
	dfs = func() bool {
		if len(res) == len(tickets)+1 {
			return true
		}
		last := res[len(res)-1]
		if _, ok := ticketMap[last]; !ok {
			return false
		}
		sortedKeys := sortKey(ticketMap[last])
		for _, next := range sortedKeys {
			if ticketMap[last][next] == 0 {
				continue
			}
			res = append(res, next)
			ticketMap[last][next]--
			if dfs() {
				return true
			}
			res = res[:len(res)-1]
			ticketMap[last][next]++
		}
		return false
	}

	has := dfs()
	if has {
		return res
	}
	return []string{}
}

func sortKey(m map[string]int) []string {
	keys := []string{}
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

// @lc code=end
