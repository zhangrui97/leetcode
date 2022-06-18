func closedIsland(grid [][]int) int {
  m, n := len(grid), len(grid[0])
  count := 0
  var dfs func(i, j int)
  dfs = func(i, j int) {
    if grid[i][j] == 1 { return }
    grid[i][j] = 1
    if i > 0 { dfs(i-1, j) }
    if i < m-1 { dfs(i+1, j) }
    if j > 0 { dfs(i, j-1) }
    if j < n-1 { dfs(i, j+1) }
  }
  for i := 0; i < m; i++ {
    dfs(i, 0)
    dfs(i, n-1)
  }
  for i := 1; i < n-1; i++ {
    dfs(0, i)
    dfs(m-1, i)
  }
  for i := 1; i < m-1; i++ {
    for j := 1; j < n-1; j++ {
      if grid[i][j] == 0 {
        count ++ 
        dfs(i, j)
      }
    }
  }
  return count
}