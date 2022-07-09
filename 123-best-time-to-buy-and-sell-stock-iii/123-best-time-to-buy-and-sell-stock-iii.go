func max(a, b int)int { if a > b { return a } else { return b } }

func maxProfit(prices []int) int {
  dpFirstBuy, dpFirstSell, dpSecondBuy, dpSecondSell := -100000, -100000, -100000, -100000
  for _, p := range prices {
    dpFirstBuy = max(dpFirstBuy, -p)
    dpFirstSell = max(dpFirstBuy + p, dpFirstSell)
    dpSecondBuy = max(dpFirstSell - p, dpSecondBuy)
    dpSecondSell = max(dpSecondBuy + p, dpSecondSell)    
  }
  return dpSecondSell
}