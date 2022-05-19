func allPathsSourceTarget(graph [][]int) [][]int {
  l := len(graph)
  ans := make([][]int, 0)
  var traverse func(path []int)
  traverse = func(path []int) {
    a := path[len(path)-1]
    if a == l-1 {
      newPath := make([]int, len(path))
      copy(newPath, path)
      ans = append(ans, newPath)
    } else {
      for _, v := range graph[a] {
        traverse(append(path, v))
      }
    }
  }
  traverse([]int{0})
  return ans
}