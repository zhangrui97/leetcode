func oddCells(m int, n int, indices [][]int) int {
  mat := make([][]int, m)
  for i := range mat {
    mat[i] = make([]int, n)
  }
  for _, v := range indices {
    r, c := v[0], v[1]
    for i := 0; i < n; i++ {
      mat[r][i] += 1
    }
    for j := 0; j < m; j++ {
      mat[j][c] += 1
    }
  }
  ans := 0
  for _, r := range mat {
    for _, v := range r {
      if v % 2 == 1 {
        ans++
      }
    }
  }
  return ans
}