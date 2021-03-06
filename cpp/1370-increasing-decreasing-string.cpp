/**
 * @file 1370-increasing-decreasing-string.cpp
 * @author krizsx (cb.kris.bj@gmail.com)
 * @brief https://leetcode-cn.com/problems/increasing-decreasing-string/
 * @version 0.1
 * @date 2020-11-25 23:06:42
 * 
 * @copyright Copyright (c) 2020
 * 
 */
#include <iostream>
#include <vector>
using namespace std;

class Solution {
   public:
    /**
     * @brief 处理字符串
     *  1.从 s 中选出 最小 的字符，将它 接在 结果字符串的后面。
     *  2.从 s 剩余字符中选出 最小 的字符，且该字符比上一个添加的字符大，将它 接在 结果字符串后面。
     *  3.重复步骤 2 ，直到你没法从 s 中选择字符。
     *  4.从 s 中选出 最大 的字符，将它 接在 结果字符串的后面。
     *  5.从 s 剩余字符中选出 最大 的字符，且该字符比上一个添加的字符小，将它 接在 结果字符串后面。
     *  重复步骤 5 ，直到你没法从 s 中选择字符。
     *  重复步骤 1 到 6 ，直到 s 中所有字符都已经被选过。
     * 
     * @param s 
     * @return string 
     */
    string sortString(string s) {
        vector<int> num(26);
        for (char &ch : s) {
            num[ch - 'a']++;
        }

        string ret;
        while (ret.length() < s.length()) {
            for (int i = 0; i < 26; i++) {
                if (num[i]) {
                    ret.push_back(i + 'a');
                    num[i]--;
                }
            }
            for (int i = 25; i >= 0; i--) {
                if (num[i]) {
                    ret.push_back(i + 'a');
                    num[i]--;
                }
            }
        }
        return ret;
    }
};

int main(int argc, char **argv) {
    auto solution = new Solution();
    auto result = solution->sortString("");
    cout << "result:" << result << endl;
}