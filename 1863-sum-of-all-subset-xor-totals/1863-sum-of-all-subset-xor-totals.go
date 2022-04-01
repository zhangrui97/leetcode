func subsetXORSum(nums []int) int {
  ans := 0
  for _, v := range nums {
    ans |= v
  }
  return ans * (1 << (len(nums) - 1))
}