func combinationSum2(candidates []int, target int) [][]int {
  n := len(candidates)
  ans := [][]int{}
  sort.Ints(candidates)
  var backtrack func(path []int , i, t int)
  backtrack = func(path []int, i, t int) {
    l := len(path)
    pre := -1
    for j := i; j < n; j++ {
      v := candidates[j]
      if t < v { return }
      if t == v {
        res := make([]int, l+1)
        copy(res, path)
        res[l] = v
        ans = append(ans, res)
        return
      }
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