func largestAltitude(gain []int) int {
  ans := 0
  sum := 0
  for _, v := range gain {
    sum += v
    if sum > ans {
      ans = sum
    }
  }
  return ans
}