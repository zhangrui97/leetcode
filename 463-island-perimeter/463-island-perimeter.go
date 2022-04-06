func islandPerimeter(grid [][]int) int {
  row := grid[0]
  r, c := len(grid), len(row)
  oneCnt := row[0]
  landCnt := 0
  for i := 1; i < c; i++ {
    if row[i] == 1 {
      oneCnt ++
      landCnt += row[i-1]
    }
  }
  for i := 1; i < r; i++ {
    row = grid[i]
    if row[0] == 1 {
      oneCnt ++
      landCnt += grid[i-1][0]
    }
    for j := 1; j < c; j++ {
      if row[j] == 1 {
        oneCnt ++
        landCnt += row[j-1]
        landCnt += grid[i-1][j]
      }
    }
  }
  return oneCnt << 2 - landCnt << 1
}