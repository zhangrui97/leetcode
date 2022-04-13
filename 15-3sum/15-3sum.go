func threeSum(nums []int) [][]int {
  sort.Ints(nums)
  length := len(nums)
  if length < 3 { return [][]int{} }
  ans := make([][]int, 0, length - 2)
  ni := nums[0] - 1
  for i := 0; i < length - 2; i++ {
    for i < length - 2 && nums[i] == ni { i++ }
    ni = nums[i]
    if ni > 0 { break }
    nj := nums[i+1] - 1
    for j := i + 1; j < length - 1; j++ {
      for j < length - 1 && nums[j] == nj { j++ }
      nj = nums[j]
      target := 0 - ni - nj
      if nj > target { break }
      l, r := j+1, length - 1
      for l <= r {
        m := (l+r) / 2
        nk := nums[m]
        if nk == target {
          
          ans = append(ans, []int{ni, nj, nk})
          break
        } else if nk < target {
          l = m + 1
        } else {
          r = m - 1
        }
      }
    }
  }
  return ans
}