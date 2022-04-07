func subarraySum(nums []int, k int) int {
  res := 0
  cache := make(map[int]int, len(nums))
  cache[0] = 1
  sum := 0
  for i := 0; i < len(nums); i++ {
    sum += nums[i]
    res += cache[sum - k]
    cache[sum]++
  }
  return res
}