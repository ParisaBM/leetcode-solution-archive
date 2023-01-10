#include <vector>
#include <iostream>

using namespace std;

class Solution {
public:
    bool solveSudoku(vector<vector<char>>& board) {
        int x, y;
        bool found = false;
        for(int i = 0; i < 9; i++) {
            for(int j = 0; j < 9; j++) {
                if(board[i][j] == '.'){
                    x = i;
                    y = j;
                    found = true;
                    break;
                }
            }
            if(found) {
                break;
            }
        }
        if(!found) {
            return true;
        }
        bool invalid[9] = {false};
        for(int i = 0; i < 9; i++) {
            for(int j = 0; j < 9; j++) {
                if((i == x || j == y || i/3 == x/3 && j/3 == y/3) && board[i][j] != '.') {
                    invalid[board[i][j] - '1'] = true;
                }
            }
        }
        for(int value = 0; value < 9; value++) {
            if(!invalid[value]) {
                board[x][y] = value + '1';
                if(solveSudoku(board)) {
                    return true;
                }
            }
        }
        board[x][y] = '.';
        return false;
    }
};

int main() {
    vector<vector<char>> board = 
    {{'5','3','.','.','7','.','.','.','.'},
    {'6','.','.','1','9','5','.','.','.'},
    {'.','9','8','.','.','.','.','6','.'},
    {'8','.','.','.','6','.','.','.','3'},
    {'4','.','.','8','.','3','.','.','1'},
    {'7','.','.','.','2','.','.','.','6'},
    {'.','6','.','.','.','.','2','8','.'},
    {'.','.','.','4','1','9','.','.','5'},
    {'.','.','.','.','8','.','.','7','9'}};
    Solution solution;
    cout << solution.solveSudoku(board) << endl;
    for(int i = 0; i < 9; i++) {
        for(int j = 0; j < 9; j++) {
            cout << board[i][j];
        }
        cout << endl;
    }
}