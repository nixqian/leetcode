#include <vector>
#include <iostream>
using namespace std;

/*
 * @lc app=leetcode.cn id=4 lang=cpp
 *
 * [4] 寻找两个正序数组的中位数
 */

// @lc code=start
class Solution {
public:
    int magicalString(int n) {
        vector<int> s;
        s.push_back(1);
        s.push_back(2);
        s.push_back(2);
        int next = 1;
        int i = 2;
        while (s.size() < n)
        {
            if (s[i] == 1)
            {
                s.push_back(next);
            } else {
                s.push_back(next);
                s.push_back(next);
            }
            i++;
            next = next == 1 ? 2 : 1;
        }
        int res = 0;
        for (int i = 0; i < n; i++)
        {
            cout << s[i] << endl;
            if (s[i] == 1)
            {
                res += 1;
            } 
        }
        return res;
    }
};
// @lc code=end


int main() {
    Solution s;
    int res = s.magicalString(6);
    cout << res << endl;
    return 0;
}