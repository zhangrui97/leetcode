func findOrder(numCourses int, prerequisites [][]int) []int {
  pre := make([][]int, numCourses)
  for i := 0; i < numCourses; i++ {
    pre[i] = []int{}
  }
  for _, v := range prerequisites {
    c := v[0]
    pre[c] = append(pre[c], v[1])
  }
  finished := make([]bool, numCourses)
  visited := make([]bool, numCourses)
  ans := []int{}
  var dfs func(i int)bool
  dfs = func(i int)bool {
    if finished[i] { return true }
    if visited[i] { return false }
    visited[i] = true
    for _, v := range pre[i] {
      if !dfs(v) {
        return false
      }
    }
    visited[i] = false
    finished[i] = true
    ans = append(ans, i)
    return true
  }
  for i := 0; i < numCourses; i++ {
    if !dfs(i) {
      return []int{}
    }
  }
  return ans
}