func isBipartite(graph [][]int) bool {
  l := len(graph)
  visited := make([]bool, l)
  flags := make([]bool, l)
  var dfs func(g int, flag bool) bool
  dfs = func(g int, flag bool) bool {
    if visited[g] { 
      return flags[g] == flag
    }
    visited[g] = true
    flags[g] = flag
    for _, v := range graph[g] {
      if !dfs(v, !flag) {
        return false
      }
    }
    return true
  }
  for i := range graph {
    for j := range flags {
      visited[j] = false
      flags[j] = false
    }
    if !dfs(i, true) {
      return false
    }
  }
  return true
}