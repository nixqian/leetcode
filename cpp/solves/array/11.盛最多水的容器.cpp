#include<vector>

using namespace std;

/*
 * @lc app=leetcode.cn id=11 lang=cpp
 *
 * [11] 盛最多水的容器
 */

// @lc code=start
class Solution {
public:
    int maxArea(vector<int>& height) {
        int i = 0, j = height.size()-1;
        int res = 0;
        while (i < j)
        {
            int minHeight = min(height[i], height[j]);
            int area = minHeight * (j-i);
            res = max(area, res);
            while (i <j && height[i] <= minHeight)
            {
                i++;
            }
            while (i<j && height[j] <= minHeight)
            {
                j--;
            }
        }
        return res;
    }
};
// @lc code=end

