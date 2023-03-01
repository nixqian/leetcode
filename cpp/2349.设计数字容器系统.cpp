/*
 * @lc app=leetcode.cn id=2349 lang=cpp
 *
 * [2349] 设计数字容器系统
 */
#include<map>
#include<vector>

using namespace std;

// @lc code=start
class NumberContainers {
public:
    map<int, int> nums;
    NumberContainers() {
    }
    
    void change(int index, int number) {
        nums[index] = number;
    }
    
    int find(int number) {

        for (auto i = nums.begin(); i != nums.end(); i++)
        {
            if (i->second == number)
            {
                return i->first;
            }
        }
        return -1;
    }
};

/**
 * Your NumberContainers object will be instantiated and called as such:
 * NumberContainers* obj = new NumberContainers();
 * obj->change(index,number);
 * int param_2 = obj->find(number);
 */
// @lc code=end

