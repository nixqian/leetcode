#include<string>
#include<vector>

using namespace std;
/*
 * @lc app=leetcode.cn id=14 lang=cpp
 *
 * [14] 最长公共前缀
 */

// @lc code=start
class Solution {
public:
    string longestCommonPrefix(vector<string>& strs) {
        string res = strs[0];
        for (int i = 1; i < strs.size(); i++)
        {
            string curStr = "";
            for (int j = 0; j < res.size() && j < strs[i].size() && res[j] == strs[i][j]; j++)
            {
                curStr = curStr.append(string(1, res[j]));
            }
            res = curStr;
        }
        return res;
    }
};
// @lc code=end

