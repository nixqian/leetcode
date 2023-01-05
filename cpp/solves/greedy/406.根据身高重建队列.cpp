#include<vector>
#include<algorithm>

using namespace std;
/*
 * @lc app=leetcode.cn id=406 lang=cpp
 *
 * [406] 根据身高重建队列
 */

// @lc code=start
class Solution {
public:
    vector<vector<int>> reconstructQueue(vector<vector<int>>& people) {
        vector<vector<int>> queue;
        sort(people.begin(), people.end(), [](const vector<int> a, const vector<int> b) -> bool {
                if (a[0] == b[0]) return a[1] < b[1];
                return a[0]>b[0];
            });
        for (int i = 0; i < people.size(); i++)
        {
            int pos = people[i][1];
            queue.insert(queue.begin() + pos, people[i]);
        }
        return queue;
    }
};
// @lc code=end

