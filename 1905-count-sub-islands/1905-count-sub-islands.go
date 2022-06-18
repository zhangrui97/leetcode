func countSubIslands(grid1 [][]int, grid2 [][]int) int {
  m, n := len(grid2), len(grid2[0])
  count := 0
  isSub := true
  var dfs func(i, j int)
  dfs = func(i, j int) {
    if grid2[i][j] == 0 { return }
    isSub = isSub && grid1[i][j] == 1
    grid2[i][j] = 0
    if i > 0 { dfs(i-1, j) }
    if i < m-1 { dfs(i+1, j) }
    if j > 0 { dfs(i, j-1) }
    if j < n-1 { dfs(i, j+1) }
  }
  for i, r := range grid2 {
    for j, v := range r {
      if v == 1 {
        isSub = true
        dfs(i, j)
        if isSub { count++ }
      }
    }
  }
  return count
}