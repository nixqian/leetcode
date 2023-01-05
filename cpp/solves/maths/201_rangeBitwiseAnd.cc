
/*
给你两个整数 left 和 right ，表示区间 [left, right] ，返回此区间内所有数字 按位与 的结果（包含 left 、right 端点）。

示例 1：
输入：left = 5, right = 7
输出：4

示例 2：
输入：left = 0, right = 0
输出：0

示例 3：
输入：left = 1, right = 2147483647
输出：0
 
提示：
0 <= left <= right <= 231 - 1
*/

// 直接运算超时， 需要优化，只要出现了0, 后续都是0
class Solution {
public:
    int rangeBitwiseAnd(int left, int right) {
        long res = left;
        for (long i = left+1; i <= right; i++)
        {
            res = res & i;
        }
        return int(res);
    }
};