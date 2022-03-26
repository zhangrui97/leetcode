/**
 * @param {number[][]} matrix
 * @return {number}
 */
var countSquares = function(matrix) {
  const rlen = matrix.length
  const clen = matrix[0].length
  const ones = new Set()
  for (let i = 0, acc = 0; i < rlen; i ++, acc += clen) {
    for (let j = 0; j < clen; j ++) {
      if (matrix[i][j]) ones.add(acc + j)
    }
  }
  let ans = ones.size
  let dp = new Set(ones)
  for (let i = 1; i < Math.min(rlen, clen); i ++) {
    const nextdp = new Set()
    for (const j of dp) {
      const r = ~~(j / clen)
      const c = j % clen
      if (r + i >= rlen || c + i >= clen) continue
      let allOnes = true
      for (let k = 0; k < i; k++) {
        if (!ones.has((r+k)*clen+c+i) || !ones.has((r+i)*clen+c+k)) {
          allOnes = false
          break
        }
      }
      if (allOnes && ones.has((r+i)*clen+c+i)) nextdp.add(j)
    }
    ans += nextdp.size
    dp = new Set(nextdp)
  }
  return ans
};