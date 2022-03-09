/**
 * @param {number[][]} grid
 * @return {number}
 */
var projectionArea = function(grid) {
  const col = grid[0].length
  const colMax = new Array(col).fill(0)
  let ans = 0
  for (let i = 0; i < grid.length; i ++) {
    const row = grid[i]
    ans += Math.max(...row)
    for (let j = 0; j < col; j++) {
      if (row[j] > 0) {
        ans ++
        colMax[j] = Math.max(colMax[j], row[j])
      }
    }
  }
  colMax.forEach(n => ans += n)
  return ans
};