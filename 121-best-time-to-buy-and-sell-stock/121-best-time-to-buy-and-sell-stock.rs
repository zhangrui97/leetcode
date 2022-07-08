impl Solution {
  pub fn max_profit(prices: Vec<i32>) -> i32 {
    let mut dp = prices[0];
    let mut ans = 0;
    for p in prices {
      dp = dp.min(p);
      ans = ans.max(p - dp);
    }
    return ans;
  }
}