/**
 * @param {number} n
 * @return {number}
 */
var climbStairs = function(n) {
  if (n < 4) return n
  let [dp0, dp1] = [2, 3]
  for (let i = 4; i <= n; i++) {
    [dp0, dp1] = [dp1, dp0 + dp1]
  }
  return dp1
};