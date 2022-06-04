func equationsPossible(equations []string) bool {
  parents := [26]int{}
  for i := 0; i < 26; i++ {
    parents[i] = -1
  }
  var find func(i int)int
  find = func(i int)int {
    if parents[i] == -1 {
      parents[i] = i
      return i
    }
    if i == parents[i] { return i }
    p := find(parents[i])
    parents[i] = p
    return p
  }
  union := func(i, j int) {
    ri, rj := find(i), find(j)
    if ri == rj { return }
    parents[rj] = ri
  }
  connected := func(i, j int)bool {
    fmt.Println(find(i), find(j))
    return find(i) == find(j)
  }
  inequals := [][2]int{}
  for _, v := range equations {
    if v[1] == '!' {
      inequals = append(inequals, [2]int{int(v[0])-int('a'), int(v[3])-int('a')})
    } else {
      union(int(v[0])-int('a'), int(v[3])-int('a'))
    }
  }
  for _, v := range inequals {
    if connected(v[0], v[1]) {
      return false
    }
  }
  return true
}