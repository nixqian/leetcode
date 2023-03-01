#include<vector>

using namespace std;
/*
 * @lc app=leetcode.cn id=264 lang=cpp
 *
 * [264] 丑数 II
 */


// @lc code=start
class Solution {
public:
    int nthUglyNumber(int n) {
        vector<int> nums(n);
        nums[0] = 1;
        int k2 = 0, k3=0, k5=0;
        for (int i = 1; i < n; i++)
        {
            int next = min(min(nums[k2] * 2, nums[k3]*3), nums[k5]*5);
            if (next == nums[k2] * 2) k2++;
            if (next == nums[k3] * 3) k3++;
            if (next == nums[k5] * 5) k5++;
            nums[i] = next;
        }
        return nums[n-1];
    }
};
// @lc code=end

