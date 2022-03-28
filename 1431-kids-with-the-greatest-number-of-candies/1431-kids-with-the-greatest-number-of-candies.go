func kidsWithCandies(candies []int, extraCandies int) []bool {
  max := 1
  for _, v := range candies {
    if v > max {
      max = v
    }
  }
  ans := make([]bool, len(candies))
  for i, v := range candies {
    ans[i] = (v+extraCandies) >= max
  }
  return ans
}