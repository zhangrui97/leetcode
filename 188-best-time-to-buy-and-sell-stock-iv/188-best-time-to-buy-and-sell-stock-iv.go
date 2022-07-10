func max(a, b int)int { if a > b { return a } else { return b } }

func maxProfit(k int, prices []int) int {
  dpBuy := make([]int, k+1)
  dpSell := make([]int, k+1)
  for i := range dpBuy {
    dpBuy[i] = -1000
  }
  for _, p := range prices {
    for i := 0; i < k; i++ {
      dpBuy[i+1] = max(dpBuy[i+1], dpSell[i] - p)
      dpSell[i+1] = max(dpSell[i+1], dpBuy[i+1] + p)
    }
  }
  return dpSell[k]
}