func max(a, b int)int { if a > b { return a } else { return b } }
func min(a, b int)int { if a > b { return b } else { return a } }

func minPathSum(grid [][]int) int {
  m, n := len(grid), len(grid[0])
  for i := n-2; i >= 0; i-- {
    grid[m-1][i] += grid[m-1][i+1]
  }
  for i := m-2; i >= 0; i-- {
    grid[i][n-1] += grid[i+1][n-1]
    for j := n-2; j >= 0; j-- {
      grid[i][j] += min(grid[i+1][j], grid[i][j+1])
    }
  }
  return grid[0][0]
}