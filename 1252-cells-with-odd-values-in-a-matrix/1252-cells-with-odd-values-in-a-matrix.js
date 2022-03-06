/**
 * @param {number} m
 * @param {number} n
 * @param {number[][]} indices
 * @return {number}
 */
var oddCells = function(m, n, indices) {
  const mat = Array(m).fill().map(_ => Array(n).fill(0))
  indices.forEach(([r, c]) => {
    for (let i = 0; i < n; i++) mat[r][i] += 1
    for (let i = 0; i < m; i++) mat[i][c] += 1
  })
  let ans = 0
  mat.forEach(r => r.forEach(c => ans += c % 2))
  return ans
};