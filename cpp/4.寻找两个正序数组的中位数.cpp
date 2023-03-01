#include<vector>
#include<iostream>

using namespace std;
/*
 * @lc app=leetcode.cn id=4 lang=cpp
 *
 * [4] 寻找两个正序数组的中位数
 */

// @lc code=start
class Solution {
public:
    double findMedianSortedArrays(vector<int>& nums1, vector<int>& nums2) {
        vector<int> nums;
        int i = 0, j = 0;
        while (i < nums1.size() && j < nums2.size())
        {
            if (nums1[i] <= nums2[j])
            {
                nums.push_back(nums1[i]);
                i++;
            } else {
                nums.push_back(nums2[j]);
                j++;
            }
        }
        while (i < nums1.size())
        {
            nums.push_back(nums1[i]);
            i++;
        }
        while (j < nums2.size()) 
        {
            nums.push_back(nums2[j]);
            j++;
        }
        int total = nums.size();
        cout<<total<<endl;
        if (total % 2 == 1)
        {
            return double(nums[total / 2]);
        } else {
            return double(nums[total / 2 - 1] + nums[total / 2]) / 2;
        }
    }
};
// @lc code=end


int main() {
    Solution s;
    vector<int> a = {1,3};
    vector<int> b = {2};
    double ret = s.findMedianSortedArrays(a, b);
    cout<<ret<<endl;
}