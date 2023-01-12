#include <iostream>
#include <map>
#include <tuple>
using namespace std;

class Solution {
    map<tuple<string, string>, bool> solutionCache;
public:
    bool isScramble(string s1, string s2) {
        if(solutionCache.count(tuple(s1, s2))) {
            return solutionCache[tuple(s1, s2)];
        }
        int n = s1.length();
        if(s1 == s2) {
            solutionCache[tuple(s1, s2)] = true;
            return true;
        }
        int charDif[26] = {0};
        for(int i = 0; i < n; i++) {
            charDif[s1[i] - 'a']++; 
        }
        for(int i = 0; i < n; i++) {
            if(--charDif[s2[i] - 'a'] < 0) {
                solutionCache[tuple(s1, s2)] = false;
                return false;
            } 
        }
        for(int i = 1; i <= n - 1; i++) {
            if(isScramble(s1.substr(0, i), s2.substr(0, i)) && isScramble(s1.substr(i), s2.substr(i)) ||
            isScramble(s1.substr(i), s2.substr(0, n-i)) && isScramble(s1.substr(0, i), s2.substr(n-i))
            ) {
                solutionCache[tuple(s1, s2)] = true;
                return true;
            }
        }
        solutionCache[tuple(s1, s2)] = false;
        return false;
    }
};

int main() {
    Solution solution;
    cout << solution.isScramble("a", "a") << endl;
    cout << solution.isScramble("great", "rgeat") << endl;
    cout << solution.isScramble("abcde", "caebd") << endl;
    cout << solution.isScramble("abc", "bca") << endl;
    cout << solution.isScramble("ymjmfxshglxwrrgufcvvzjuietjzzz", "fxczujvmwizrzgxgjmvzelyjthusrf") << endl;
    cout << solution.isScramble("abcdefghijklmnopqrstuvwxyz", "zyxwvutsrqponmlkjihgfedcba") << endl;
    return 0;
}