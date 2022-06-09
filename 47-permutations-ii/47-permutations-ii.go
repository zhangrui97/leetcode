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
    for i, v := range nums {
      if used[i] || v == pre { continue }
      track = append(track, v)
      used[i] = true
      pre = v
      backtrack()
      track = track[:len(track)-1]
      used[i] = false
    }
  }
  backtrack()
  return result
}