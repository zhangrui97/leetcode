func generate(numRows int) [][]int {
  dp := make([][]int, numRows)  
  for i := 0; i < numRows; i++ {
    row := make([]int, i + 1)
    row[0] = 1
    row[i] = 1
    for j := 1; j <= i/2; j++ {
      row[j] = dp[i-1][j-1] + dp[i-1][j]
      row[i-j] = row[j]
    }
    dp[i] = row
  }
  return dp
}