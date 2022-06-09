func combinationSum(candidates []int, target int) [][]int {
  n := len(candidates)
  ans := make([][]int, 0)
  sort.Ints(candidates)
  var backtrack func(path []int, i, target int)
  backtrack = func(path []int, i, target int) {
    l := len(path)
    if target == 0 {
      res := make([]int, l)
      copy(res, path)
      ans = append(ans, res)
      return
    }
    for ; i < n; i++ {
      v := candidates[i]
      if target < v { break }
      path = append(path, v)
      backtrack(path, i, target - v)
      path = path[:l]
    }
  }
  backtrack([]int{}, 0, target)
  return ans
}