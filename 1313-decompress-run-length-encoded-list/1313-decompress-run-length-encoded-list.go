func decompressRLElist(nums []int) []int {
  var ans []int
  for i := 0; i < len(nums); i += 2 {
    v := nums[i+1]
    for j := 0; j < nums[i]; j++ {
      ans = append(ans, v)
    }
  }
  return ans
}