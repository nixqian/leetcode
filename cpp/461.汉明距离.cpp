/*
 * @lc app=leetcode.cn id=461 lang=cpp
 *
 * [461] 汉明距离
 */

// @lc code=start
class Solution {
public:
    int hammingDistance(int x, int y) {
        int res = 0;
        while (x != 0 || y != 0)
        {
            int resx = x % 2;
            int resy = y % 2;
            if (resx != resy)
            {
                res++;
            }
            x = x / 2;
            y = y / 2;
        }
        return res;
    }
};
// @lc code=end

