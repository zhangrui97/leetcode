func combinationSum2(candidates []int, target int) [][]int {
  n := len(candidates)
  ans := [][]int{}
  sort.Ints(candidates)
  var backtrack func(path []int , i, t int)
  backtrack = func(path []int, i, t int) {
    l := len(path)
    if t == 0 {
      res := make([]int, l)
      copy(res, path)
      ans = append(ans, res)
      return
    }
    pre := -1
    for j := i; j < n; j++ {
      v := candidates[j]
      if t < v { return }
      if v == pre { continue }
      path = append(path, v)
      pre = v
      backtrack(path, j+1, t-v)
      path = path[:l]
    }
  }
  backtrack([]int{}, 0, target)
  return ans
}