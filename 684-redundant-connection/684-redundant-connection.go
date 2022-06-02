func findRedundantConnection(edges [][]int) []int {
  n := len(edges)
  parents := make([]int, n+1)
  for i := 1; i <= n; i++ {
    parents[i] = i
  }
  var find func(i int)int
  find = func(i int)int {
    if parents[i] == i {
      return i
    }
    p := find(parents[i])
    parents[i] = p
    return p
  }
  union := func(i, j int)bool {
    ri, rj := find(i), find(j)
    if ri == rj {
      return false
    }
    parents[rj] = parents[ri]
    return true
  }
  for _, v := range edges {
    if !union(v[0], v[1]) {
      return v
    }
  }
  return []int{}
}