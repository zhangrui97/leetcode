func maximumWealth(accounts [][]int) int {
  ans := 0
  for _, r := range accounts {
    sum := 0
    for _, v := range r {
      sum += v
    }
    if sum > ans {
      ans = sum
    }
  }
  return ans
}