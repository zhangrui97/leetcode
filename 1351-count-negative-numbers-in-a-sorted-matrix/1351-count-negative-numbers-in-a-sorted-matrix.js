/**
 * @param {number[][]} grid
 * @return {number}
 */
var countNegatives = function(grid) {
  const m = grid.length
  const n = grid[0].length
  let ans = 0
  for (let i = 0; i < m; i++) {
    const index = grid[i].findIndex(el => el < 0)
    if (index >= 0) ans += n - index    
  }
  return ans
};