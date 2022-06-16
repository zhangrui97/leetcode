func canPartitionKSubsets(nums []int, k int) bool {
  if k == 1 { return true }
  n := len(nums)
  if k > n {
    return false
  }
  sum := 0
  sort.Ints(nums)
  for _, v := range nums {
    sum += v
  }
  if sum % k != 0 { return false }
  mean := sum / k
  if nums[n-1] > mean { return false }
  used := 0
  memo := make(map[int]bool)
  var backtrack func(l, start, t int)bool
  backtrack = func(l, start, t int)bool {
    if l == k { return true }
    if t == 0 {
      res := backtrack(l+1, 0, mean)
      memo[used] = res
      return res
    }
    if res, ok := memo[used]; ok {
      return res
    }
    pre := 0
    for j := start; j < n; j++ {
      if used >> j & 1 == 1 { continue }
      v := nums[j]
      if t < v { return false }
      if v == pre { continue }
      used |= 1 << j
      pre = v
      if backtrack(l, j + 1, t - v) {
        return true
      }
      used &= ^(1 << j)
    }
    return false
  }
  return backtrack(0, 0, mean)
}