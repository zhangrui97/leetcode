func diagonalSum(mat [][]int) int {
  n := len(mat)
  ans := 0
  for i := 0; i < n/2; i++ {
    ans += mat[i][i] + mat[i][n-1-i] + mat[n-1-i][n-1-i] + mat[n-1-i][i]
  }
  if n % 2 == 1 {
    ans += mat[n/2][n/2]
  }
  return ans
}