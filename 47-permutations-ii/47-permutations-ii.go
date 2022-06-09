func permuteUnique(nums []int) [][]int {
  n := len(nums)
  if n == 1 { return [][]int{nums} }
  sort.Ints(nums)
  result := make([][]int, 0)
  track := make([]int, 0, n)
  used := make([]bool, n)
  var backtrack func()
  backtrack = func() {
    if len(track) == n {
      res := make([]int, n)
      copy(res, track)
      result = append(result, res)
      return
    }
    pre := -100
    for i := 0; i < n; i++ {
      if used[i] || nums[i] == pre { continue }
      track = append(track, nums[i])
      used[i] = true
      pre = nums[i]
      backtrack()
      track = track[:len(track)-1]
      used[i] = false
    }
  }
  backtrack()
  return result
}