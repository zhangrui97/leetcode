func coinChange(coins []int, amount int) int {
  if amount == 0 { return 0 }
  ans := make([]int, amount+1)
  min := 13
  for _, v := range coins {
    if v <= amount {
      ans[v] = 1
      if v == amount {
        return 1
      }
    }
    if v < min {
      min = v
    }
  }

  for i := min+1; i <= amount; i++ {
    if ans[i] > 0 { continue }
    min = amount+1
    for _, v := range coins {
      if v < i && ans[i-v] > 0 {
        if ans[i-v] < min {
          min = ans[i-v]
        }
      }
    }
    if min != amount+1 {
      ans[i] = min + 1
    }
  }
  if ans[amount] == 0 {
    return -1
  }
  return ans[amount]
}