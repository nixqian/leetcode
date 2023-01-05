#include<vector>

using namespace std;
/*
 * @lc app=leetcode.cn id=526 lang=cpp
 *
 * [526] 优美的排列
 */

// @lc code=start
class Solution {
public:
    int res = 0;
    vector<int> used = vector<int>(20, 0);

    int countArrangement(int n) {
        backtrack(n, 1);
        return res;
    }

    void backtrack(int n, int idx) {
        if (idx > n)
        {
            res++;
            return;
        }
        for (int num = 1; num <= n; num++)
        {
            if (used[num] == 1)
            {
                continue;
            }
            if (num % idx != 0 && idx % num != 0)
            {   
                continue;
            }
            used[num]++;
            backtrack(n, idx+1);
            used[num]--;
        }
        return;
    }
};
// @lc code=end

