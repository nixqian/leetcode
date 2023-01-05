#include<vector>
#include<unordered_map>

using namespace std;
/*
 * @lc app=leetcode.cn id=137 lang=cpp
 *
 * [137] 只出现一次的数字 II
 */

// @lc code=start
class Solution {
public:
    int singleNumber(vector<int>& nums) {
        unordered_map<int, int> umap;
        for (int num : nums)
        {
            umap[num]++;
        }
        for (auto [num, count] : umap) {
            if (count == 1)
            {
                return num;
            }
        }
        return 0;
    }
};
// @lc code=end

