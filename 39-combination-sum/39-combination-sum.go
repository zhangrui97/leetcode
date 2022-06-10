func combinationSum(candidates []int, target int) [][]int {
  n := len(candidates)
  ans := make([][]int, 0)
  sort.Ints(candidates)
  var backtrack func(path []int, i, target int)
  backtrack = func(path []int, i, target int) {
    l := len(path)
    for ; i < n; i++ {
      v := candidates[i]
      if target < v { return }
      if target == v {
        res := make([]int, l+1)
        copy(res, path)
        res[l] = v
        ans = append(ans, res)
        return
      }
      path = append(path, v)
      backtrack(path, i, target - v)
      path = path[:l]
    }
  }
  backtrack([]int{}, 0, target)
  return ans
}