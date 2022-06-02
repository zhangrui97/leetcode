func findCircleNum(isConnected [][]int) int {
  n := len(isConnected)
  parents := make([]int, n)
  root := func(i int)int {
    for i != parents[i] {
      parents[i] = parents[parents[i]]
      i = parents[i]
    }
    return i
  }
  for i := 0; i < n; i++ {
    parents[i] = i
  }
  count := n
  if n == 1 { return 1 }
  for i := 0; i < n; i++ {
    for j := i+1; j < n; j++ {
      if isConnected[i][j] == 1 {
        ri, rj := root(i), root(j)
        if ri != rj {
          count--
        }
        parents[rj] = parents[ri]
      }
    }
  }
  return count
}