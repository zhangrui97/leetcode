func combine(n int, k int) [][]int {
  ans := [][]int{}
  var backtrack func(path []int, i, l int)
  backtrack = func(path []int, i, l int) {
    if l == k {
      res := make([]int, k)
      copy(res, path)
      ans = append(ans, res)
      return
    }
    for j := i+1; j <= n; j++ {
      path = append(path, j)
      backtrack(path, j, l+1)
      path = path[:l]
    }
  }
  backtrack([]int{}, 0, 0)
  return ans
}