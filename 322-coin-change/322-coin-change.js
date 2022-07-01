/**
 * @param {number[]} coins
 * @param {number} amount
 * @return {number}
 */
var coinChange = function(coins, amount) {
  const dp = Array.apply(null, Array(amount+1)).map(_ => Infinity)
  dp[0] = 0;
  for (i = 1; i <= amount; i++) {
    for (let coin of coins) {
      if(i-coin >= 0) { 
              
        dp[i] = Math.min(dp[i], dp[i-coin]+1);
      }
    }
       
  }
  if (dp[amount] == Infinity) {
    return -1
  } 
  return dp[amount];
};