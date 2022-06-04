func diff(a, b int)int {
  if a > b {
    return a - b
  } else {
    return b - a
  }
}

func dist(points [][]int, i, j int)int {
  return diff(points[i][0], points[j][0]) + diff(points[i][1], points[j][1])
}

type Edge struct {
  pair int
  length int
}

func minCostConnectPoints(points [][]int) int {
  l := len(points)
  if l == 1 { return 0 }
  if l == 2 { return dist(points, 0, 1) }
  g := make([]int, l)
  var find func(i int)int
  find = func(i int)int {
    if g[i] == i { return i }
    r := find(g[i])
    g[i] = r
    return r
  }
  count := l
  union := func(pair int) {
    ri, rj := find(pair >> 16), find(pair & 0xFFFF)
    if ri == rj { return }
    g[rj] = ri
    count--
  }
  connected := func(pair int)bool {
    return find(pair >> 16) == find(pair & 0xFFFF)
  }
  
  e := make([]Edge, 0, l*(l-1)/2)  
  for i := 0; i < l; i ++ {
    g[i] = i
    for j := i+1; j < l; j++ {
      e = append(e, Edge{i << 16 | j, dist(points, i, j)})
    }
  }
  sort.Slice(e, func(i, j int)bool { return e[i].length < e[j].length })
  res := 0
  for _, v := range e {
    if connected(v.pair) { continue }
    union(v.pair)
    res += v.length
    if count == 1 { break }  
  }
  return res
}