#include<unordered_map>
#include<string>
#include<iostream>

using namespace std;
/*
 * @lc app=leetcode.cn id=1797 lang=cpp
 *
 * [1797] 设计一个验证系统
 */

// @lc code=start
class AuthenticationManager {
public:
    int ttl;
    unordered_map<string, int> timeoutMap;
    AuthenticationManager(int timeToLive) {
        ttl = timeToLive;
        // cout<<"ttl: "<<ttl<<endl;
    }
    
    void generate(string tokenId, int currentTime) {
        timeoutMap[tokenId] = currentTime + this->ttl;
        // cout<< currentTime <<" generate: "<<tokenId<<": "<<timeoutMap[tokenId]<<endl;
        return;
    }
    
    void renew(string tokenId, int currentTime) {
        if (this->timeoutMap.find(tokenId) != this->timeoutMap.end() && this->timeoutMap[tokenId] > currentTime)
        {
            this->timeoutMap[tokenId] = currentTime + this->ttl;
        } 
        // cout<< currentTime <<" renew: "<<tokenId<<": "<<timeoutMap[tokenId]<<endl;
    }
    
    int countUnexpiredTokens(int currentTime) {
        int count = 0;
        cout<< currentTime<<" count: ";
        for (auto i = this->timeoutMap.begin(); i != this->timeoutMap.end(); i++)
        {
            // cout << "(" << i->first << ":" << i->second<<")";
            if (i->second > currentTime)
            {
                count++;
            } 
        }
        // cout << "total: "<< count<< endl;
        return count;
    }
};

/**
 * Your AuthenticationManager object will be instantiated and called as such:
 * AuthenticationManager* obj = new AuthenticationManager(timeToLive);
 * obj->generate(tokenId,currentTime);
 * obj->renew(tokenId,currentTime);
 * int param_3 = obj->countUnexpiredTokens(currentTime);
 */
// @lc code=end

