func max(a, b int)int { if a > b { return a } else { return b} }

func maxProfit(prices []int) int {
  n := len(prices)
  if n == 1 { return 0 }
  if n == 2 { return max(0, prices[1] - prices[0]) }
  buy := make([]int, n)
  buy[0] = -prices[0]
  buy[1] = max(buy[0], -prices[1])
  for i := 2; i < n; i++ { buy[i] = -1000 }
  sell := make([]int, n-1)
  for i := 2; i < n; i++ {
    sell[i-1] = max(sell[i-2], buy[i-2]+prices[i-1])
    buy[i] = max(buy[i-1], sell[i-2]-prices[i])
  }
  fmt.Println(buy, sell, buy[n-2]+prices[n-1])
  return max(sell[n-2], buy[n-2]+prices[n-1])
}