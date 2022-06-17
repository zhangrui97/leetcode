func subsets(nums []int) [][]int {
  n := len(nums)
  ans := [][]int{}
  var backtrack func(path int, i int)
  backtrack = func(path int, i int) {
    res := []int{}
    for j, p := 0, path; p > 0; j, p = j+1, p>>1 {
      if (p & 1) > 0 { res = append(res, nums[j]) }
    }
    ans = append(ans, res)
    for j := i; j < n; j++ {
      backtrack(path | (1 << j), j+1)
    }
  }
  backtrack(0, 0)
  return ans
}