func fourSum(nums []int, target int) [][]int {
  length := len(nums)
  if length < 4 { return [][]int{} }
  avg := int(math.Ceil(float64(target >> 2)))
  avg2 := int(math.Ceil(float64(target >> 1)))
  sort.Ints(nums)
  ans := make([][]int, 0, length - 2)
  ai := nums[0] - 1
  for i := 0; i < length - 3; i++ {
    for ; i < length - 3 && nums[i] == ai; i++ {}
    if i >= length - 3 { break }
    ai = nums[i]
    if ai > avg { break }
    aj := nums[i+1] - 1
    for j := i + 1; j < length - 2; j++ {
      for ; j < length - 2 && nums[j] == aj; j++ {}
      if j >= length - 2 { break }
      aj = nums[j]
      sum := ai + aj
      if sum > avg2 { break }
      t := target - sum
      l, r := j+1, length - 1
      for l < r {
        al := nums[l]
        ar := nums[r]
        if al + ar == t {
          ans = append(ans, []int{ai, aj, al, ar})
          for l++; l < r && nums[l] == al; l++ {}
          for r--; l < r && nums[r] == ar; r-- {}
        } else if al + ar > t {
          for r--; l < r && nums[r] == ar; r-- {}
        } else {
          for l++; l < r && nums[l] == al; l++ {}
        }
      }
    }
  }
  return ans
}