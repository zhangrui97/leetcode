public class Solution {
  private bool isValid(char[][] b, int i, int j, char n) {
    foreach (char ch in b[i]) {
      if (ch == n) return false;
    }
    for (int k = 0; k < 9; k++) {
      if (b[k][j] == n) return false;
    }
    int fr = i / 3 * 3;
    int fc = j / 3 * 3;
    for (int r = fr; r < fr + 3; r++) {
      for (int c = fc; c < fc + 3; c++) {
        if (b[r][c] == n) return false;
      }
    }
    return true;
  }
  private bool backtrack(char[][] b, int i, int j) {
    if (j == 9) {
      return backtrack(b, i+1, 0);
    }
    if (i == 9) {
      return true;
    }
    if (b[i][j] != '.') {
      return backtrack(b, i, j+1);
    }
    for (char ch = '1'; ch <= '9'; ch++) {
      if (isValid(b, i, j, ch)) {
        b[i][j] = ch;
        if (backtrack(b, i, j+1))
          return true;
        b[i][j] = '.';
      }
    }
    return false;
  }
  public void SolveSudoku(char[][] board) {
    backtrack(board, 0, 0);
  }
}