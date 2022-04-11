func lengthOfLIS(nums []int) int {
  l := len(nums)
  heap := make([]int, l)
  heap[0] = nums[0]
  result := 1
  for i := 1; i < l; i++ {
    target := nums[i]
    l, r := 0, result
    for l < r {
      m := (l + r) / 2
      if heap[m] < target {
        l = m + 1
      } else {
        r = m
      }
    }  
    if l >= result {
      result++
    }
    heap[l] = target
  }
  return result
}