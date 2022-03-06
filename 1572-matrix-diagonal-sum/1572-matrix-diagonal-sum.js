/**
 * @param {number[][]} mat
 * @return {number}
 */
var diagonalSum = function(mat) {
  const n = mat.length
  if (n === 1) return mat[0][0]
  const ni = n - 1
  let ans = 0
  const half = ~~(n/2)
  for (let i = 0; i < half; i++) {
    ans += mat[i][i] + mat[i][ni-i] + mat[ni-i][i] + mat[ni-i][ni-i]
  }
  if (n % 2) {
    ans += mat[half][half]
  }
  return ans
};