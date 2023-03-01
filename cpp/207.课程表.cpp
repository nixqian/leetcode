#include<vector>

using namespace std;
/*
 * @lc app=leetcode.cn id=207 lang=cpp
 *
 * [207] 课程表
 */

// @lc code=start
class Solution {
public:
    bool canFinish(int numCourses, vector<vector<int>>& prerequisites) {
        vector<int> courseDepend;
        for (int i = 0; i < numCourses; i++)
        {
            courseDepend.push_back(-1);
        }
        for (int i = 0; i < prerequisites.size(); i++)
        {
            courseDepend[prerequisites[i][0]] = prerequisites[i][1];
        }
    }
};
// @lc code=end

