/**
 * @param {number[][]} matrix
 * @return {number[]}
 */
var luckyNumbers  = function(matrix) {
  const ans = []
  const rs = matrix.length
  matrix.forEach(row =>{
    const min = Math.min(...row)
    const j = row.indexOf(min)
    for (let i = 0; i < rs; i++) {
      if (matrix[i][j] > min) return
    }
    ans.push(min)
  })
  return ans
};