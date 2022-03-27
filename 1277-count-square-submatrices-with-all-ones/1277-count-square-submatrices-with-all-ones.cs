public class Solution {
  public int CountSquares(int[][] matrix) {
    int rlen = matrix.Length;
    int clen = matrix[0].Length;
    
    int count = 0;
    for (int i = 0; i < clen; i++) {
      if (matrix[0][i] == 1) count++;
    }
    for (int i = 1; i < rlen; i++) {
      if (matrix[i][0] == 1) count++;
    }
    for (int i = 1; i < rlen; i++) {
      for (int j = 1; j < clen; j++) {
        if (matrix[i][j] == 1) {
          count++;
          int min = Math.Min(Math.Min(matrix[i-1][j-1], matrix[i-1][j]), matrix[i][j-1]);
          count += min;
          matrix[i][j] = min + 1;
        }
      }
    }
    return count;
  }
}