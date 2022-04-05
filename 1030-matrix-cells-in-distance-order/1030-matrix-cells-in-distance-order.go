func allCellsDistOrder(rows int, cols int, rCenter int, cCenter int) [][]int {
  dist := func (r, c int) int {
    dr, dc := r - rCenter, c - cCenter
    if dr < 0 { dr = -dr }
    if dc < 0 { dc = -dc }
    return dr + dc
  }
  idx := make([][]int, 0, rows * cols)
  res := make([][]int, 0, rows * cols)
  index := 0
  for i := 0; i < rows; i++ {
    for j := 0; j < cols; j++ {
      idx = append(idx, []int{index, dist(i, j)})
      res = append(res, []int{i, j})
      index++
    }
  }
  sort.Slice(idx, func (i, j int) bool {
    return idx[i][1] < idx[j][1]
  })
  ans := make([][]int, rows * cols)
  for i := 0; i < rows*cols; i++ {
    ans[i] = res[idx[i][0]]
  }
  return ans
}