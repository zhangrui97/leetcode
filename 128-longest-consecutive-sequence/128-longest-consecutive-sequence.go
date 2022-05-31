func longestConsecutive(nums []int) int {
  l := len(nums)
  if l < 2 { return l }
  sort.Ints(nums)
  count := 1
  lo := nums[0]
  temp := 1
  for _, v := range nums {
    if v == lo + 1 {
      lo++
      temp++
      if temp > count {
        count = temp
      }
    } else if v > lo {
      lo = v
      temp = 1
    }
  }
  return count
}