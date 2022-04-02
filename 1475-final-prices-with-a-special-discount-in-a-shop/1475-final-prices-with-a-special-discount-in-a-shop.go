func finalPrices(prices []int) []int {
  for i := 0; i < len(prices); i ++ {
    p := prices[i]
    for j := i + 1; j < len(prices); j ++ {
      if prices[j] <= p {
        prices[i] -= prices[j]
        break
      }
    }
  }
  return prices  
}