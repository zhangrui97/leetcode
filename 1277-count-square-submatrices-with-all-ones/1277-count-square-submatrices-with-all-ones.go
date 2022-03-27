func countSquares(matrix [][]int) int {
  ans := 0
  for i, r := range matrix {
    for j, v := range r {
      if v == 1 {
        ans += 1
        if i > 0 && j > 0 {
          min := matrix[i-1][j-1]
          if min > matrix[i-1][j] {
            min = matrix[i-1][j]
          }
          if min > r[j-1] {
            min = r[j-1]
          }
          ans += min
          r[j] = min + 1
        }
      }
    }
  }
  return ans
}