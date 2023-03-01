#include<string>
using namespace std;

/*
 * @lc app=leetcode.cn id=13 lang=cpp
 *
 * [13] 罗马数字转整数
 */

// @lc code=start
class Solution {
public:
    int romanToInt(string s) {
        int res = 0;
        for (int i = 0; i < s.size(); i++)
        {
            if (i+1 < s.size() && romanMap(s.substr(i, 2)) != 0)
            {
                res += romanMap(s.substr(i, 2));
                i++;
            } else {
                res += romanMap(s.substr(i, 1));
            }
        }
        return res;
    }

    int romanMap(string s) {
        if (s == "I")
        {
            return 1;
        } else if (s == "V")
        {   
            return 5;
        } else if (s == "X")
        {
            return 10;
        } else if (s == "L")
        {
            return 50;
        } else if (s == "C")
        {
            return 100;
        } else if (s == "D")
        {
            return 500;
        } else if (s == "M")
        {
            return 1000;
        } else if (s == "IV")
        {
            return 4;
        } else if (s == "IX")
        {
            return 9;
        } else if (s == "XL")
        {
            return 40;
        } else if (s == "XC")
        {
            return 90;
        } else if (s == "CD")
        {
            return 400;
        } else if (s == "CM")
        {
            return 900;
        }
        return 0;
    }
};
// @lc code=end

    