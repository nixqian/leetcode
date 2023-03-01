#include<vector>
#include<algorithm>

using namespace std;
/*
 * @lc app=leetcode.cn id=2148 lang=cpp
 *
 * [2148] 元素计数
 */

// @lc code=start
class Solution {
public:
    int countElements(vector<int>& nums) {
        int count = 0;
        sort(nums.begin(), nums.end());
        for (int i = 1; i < nums.size()-1; i++)
        {
            if (nums[i]>nums[0] && nums[i]<nums[nums.size()-1])
            {
                count++;
            }
        }
        return count;
    }
};
// @lc code=end

