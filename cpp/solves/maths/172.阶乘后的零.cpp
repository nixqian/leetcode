/*
 * @lc app=leetcode.cn id=172 lang=cpp
 *
 * [172] 阶乘后的零
 */

// @lc code=start
class Solution {
public:
    int trailingZeroes(int n) {
        int count = 0;
        int x = 5;
        while (n >= x)
        {
            count += (n / x);
            x = x * 5;
        }
        return count;
    }
};
// @lc code=end

