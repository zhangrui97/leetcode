/**
 * @param {number[]} cost
 * @return {number}
 */
var minCostClimbingStairs = function(cost) {
  let [dp0, dp1] = cost
  for (let i = 2; i < cost.length; i++) {
    [dp0, dp1] = [dp1, Math.min(dp0, dp1) + cost[i]]
  }    
  return Math.min(dp0, dp1)
};