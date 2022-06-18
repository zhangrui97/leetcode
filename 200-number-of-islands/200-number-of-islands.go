func numIslands(grid [][]byte) int {
  m := len(grid)
  n := len(grid[0])
  count := 0
  var dfs func(i, j int)
  dfs = func(i, j int) {
    grid[i][j] = '0'
    if i > 0 && grid[i-1][j] == '1' { dfs(i-1, j) }
    if i < m-1 && grid[i+1][j] == '1' { dfs(i+1, j) }
    if j > 0 && grid[i][j-1] == '1' { dfs(i, j-1) }
    if j < n-1 && grid[i][j+1] == '1' { dfs(i, j+1) }
  }
  for i, r := range grid {
    for j, v := range r {
      if v == '1' {
        dfs(i, j)
        count ++
      }
    }
  }
  return count
}