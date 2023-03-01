/*
 * @lc app=leetcode.cn id=6 lang=cpp
 *
 * [6] N 字形变换
 */

#include<string>
#include<vector>

using namespace std;

// @lc code=start
class Solution {
public:
    string convert(string s, int numRows) {
        if (numRows == 1)
        {
            return s;
        }
        
        vector<string> rows;
        vector<int> rowIdx;
        for (int i = 0; i < numRows; i++)
        {
            rows.push_back(string(""));
            rowIdx.push_back(i);
        }
        for (int i = numRows-2; i > 0; i--)
        {
            rowIdx.push_back(i);
        }
        for (int i = 0; i < s.length(); i++)
        {
            int rowI = rowIdx[i % rowIdx.size()];
            rows[rowI] += s[i];
        }
        string res = "";
        for (int i = 0; i < rows.size(); i++)
        {
            res = res + rows[i];
        }
        return res;
    }
};
// @lc code=end

