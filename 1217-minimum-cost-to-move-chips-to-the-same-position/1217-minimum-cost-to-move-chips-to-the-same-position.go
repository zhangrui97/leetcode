func minCostToMoveChips(position []int) int {
  odd, even := 0, 0
  for _, v := range position {
    if v % 2 == 1 {
      odd++
    } else {
      even++
    }
  }
  if odd > even {
    return even
  } else {
    return odd
  }
}