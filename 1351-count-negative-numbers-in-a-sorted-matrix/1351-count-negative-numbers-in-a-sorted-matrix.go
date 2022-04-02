func findFirstNegativeIndex(row []int, l int) int {
  r := len(row) - 1
  for l <= r {
    mid := (l + r) / 2
    if row[mid] >= 0 {
      l = mid + 1
    } else {
      r = mid - 1
    }
  }
  return l
}
func countNegatives(grid [][]int) int {
  m := len(grid)
  n := len(grid[0])
  j := 0
  ans := 0
  for i := m - 1; i >= 0; i -- {
    j = findFirstNegativeIndex(grid[i], j)
    ans += n - j
  }
  return ans
}