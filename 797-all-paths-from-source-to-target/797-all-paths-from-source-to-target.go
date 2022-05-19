func allPathsSourceTarget(graph [][]int) [][]int {
  l := len(graph)
  ans := make([][]int, 0)
  var traverse func(a int, path []int)
  traverse = func(a int, path []int) {
    if a == l-1 {
      pl := len(path)
      newPath := make([]int, pl + 1)
      copy(newPath, path)
      newPath[pl] = a
      ans = append(ans, newPath)
    } else {
      for _, v := range graph[a] {
        traverse(v, append(path, a))
      }
    }
  }
  traverse(0, []int{})
  return ans
}