func subsets(nums []int) [][]int {
  n := len(nums)
  ans := [][]int{}
  var backtrack func(path []int, i int)
  backtrack = func(path []int, i int) {
    if i > n { return }
    l := len(path)
    res := make([]int, l)
    copy(res, path)
    ans = append(ans, res)
    for j := i; j < n; j++ {
      path = append(path, nums[j])
      backtrack(path, j+1)
      path = path[:l]
    }
  }
  backtrack([]int{}, 0)
  return ans
}