func sumOfUnique(nums []int) int {
  dict := make(map[int]int, len(nums))
  for _, v := range nums {
    dict[v]++
  }
  ans := 0
  for k, v := range dict {
    if v == 1 {
      ans += k
    }
  }
  return ans
}