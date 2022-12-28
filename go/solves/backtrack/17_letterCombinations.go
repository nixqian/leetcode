package backtrack

var numLetterMap = map[string][]string{
	"2": {"a", "b", "c"},
	"3": {"d", "e", "f"},
	"4": {"g", "h", "i"},
	"5": {"j", "k", "l"},
	"6": {"m", "n", "o"},
	"7": {"p", "q", "r", "s"},
	"8": {"t", "u", "v"},
	"9": {"w", "x", "y", "z"},
}

func letterCombinations(digits string) []string {
	if digits == "" {
		return nil
	}
	res := []string{}
	path := ""
	var backtrack func(digits string, idx int)
	backtrack = func(digits string, idx int) {
		if len(path) == len(digits) {
			res = append(res, path)
			return
		}
		cur := string(digits[idx])
		letters := numLetterMap[cur]
		for _, c := range letters {
			path = path + c
			backtrack(digits, idx+1)
			path = path[:len(path)-1]
		}
	}
	backtrack(digits, 0)
	return res
}
