#include <vector>
#include <deque>
using namespace std;

/* 

输入数组 nums 和 k，表示你每次最多能走k步，比如从0跳到k。落点的nums[k] 就是获得的分
你需要到达最后一个点，计算所经过的点的最大分值

*/
class Solution {
public:
    int maxResult(vector<int>& nums, int k) {
        deque<pair<int, int>> dQue; // idx, dp[idx]
        dQue.push_back(pair<int, int>(0, nums[0]));
        for (int i = 1; i < nums.size(); i++)
        {
            while (i - dQue.front().first > k)
            {
                dQue.pop_front();
            }
            int res = dQue.back().second + nums[i];
            while (!dQue.empty() && dQue.back().second<res)
            {
                dQue.pop_back();
            }
            dQue.push_back(pair<int, int>(i, res));
        }
        return dQue.back().second;
    }
};