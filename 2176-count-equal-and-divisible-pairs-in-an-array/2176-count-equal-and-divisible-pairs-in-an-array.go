func countPairs(nums []int, k int) int {
  ans := 0
  for i := 0; i < len(nums) - 1; i++ {
    vi := nums[i]
    for j:= i+1; j < len(nums); j++ {
      if vi == nums[j] && (i*j)%k == 0 {
        ans++
      }
    }
  }
  return ans
}