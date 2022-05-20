func canFinish(numCourses int, prerequisites [][]int) bool {
  req := make([][]int, numCourses)
  for i := range req {
    req[i] = []int{}
  }
  for _, v := range prerequisites {
    req[v[0]] = append(req[v[0]], v[1])
  }
  finish := make([]bool, numCourses)
  visit := make([]bool, numCourses)
  var dfs func(i int)bool
  dfs = func(i int)bool {
    if finish[i] { return true }
    if visit[i] { return false }
    visit[i] = true
    for _, v := range req[i] {
      if !(dfs(v)) {
        return false
      }
    }
    visit[i] = false
    finish[i] = true
    return true
  }
  for i := 0; i < numCourses; i++ {
    if !(dfs(i)) {
      return false
    }
  }
  return true
}