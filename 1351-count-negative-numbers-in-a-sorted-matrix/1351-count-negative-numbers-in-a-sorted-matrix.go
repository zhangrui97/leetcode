func countNegatives(grid [][]int) int {
  m := len(grid)
  n := len(grid[0])
  var j int
  ans := 0
  for i := m - 1; i >= 0; i -- {
    row := grid[i]
    for ; j < n; j ++ {
      if row[j] < 0 {
        break
      }
    }
    ans += n - j
  }
  return ans
}