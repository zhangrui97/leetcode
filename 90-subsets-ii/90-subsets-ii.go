func subsetsWithDup(nums []int) [][]int {
  n := len(nums)
  sort.Ints(nums)
  result := [][]int{}
  var dfs func(path, i int)
  dfs = func(path, i int) {
    res := []int{}
    for j, p := 0, path; p > 0; j, p = j+1, p >> 1 {
      if (p & 1 > 0) { res = append(res, nums[j]) }
    }
    result = append(result, res)
    pre := -20
    for j := i; j < n; j++ {
      v := nums[j]
      if v == pre { continue }
      pre = v
      dfs(path | (1 << j), j+1)
    }
  }
  dfs(0, 0)
  return result
}