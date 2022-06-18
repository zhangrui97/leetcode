func maxAreaOfIsland(grid [][]int) int {
  m, n := len(grid), len(grid[0])
  max := 0
  var dfs func(i, j int)int
  dfs = func(i, j int)int {
    if grid[i][j] == 0 { return 0 }
    grid[i][j] = 0
    size := 1
    if i > 0 { size += dfs(i-1, j) }
    if i < m-1 { size += dfs(i+1, j) }
    if j > 0 { size += dfs(i, j-1) }
    if j < n-1 { size += dfs(i, j+1) }
    return size
  }
  for i, r := range grid {
    for j, v := range r {
      if v == 1 {
        size := dfs(i, j)
        if size > max { max = size }
      }
    }
  }
  return max
}