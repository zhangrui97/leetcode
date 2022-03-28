func shuffle(nums []int, n int) []int {
  ans := make([]int, n*2)
  for i := 0; i < n; i ++ {
    ans[2*i] = nums[i]
    ans[2*i+1] = nums[n+i]
  }  
  return ans
}