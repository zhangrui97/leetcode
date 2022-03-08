/**
 * @param {number[][]} mat
 * @param {number} k
 * @return {number[]}
 */
var kWeakestRows = function(mat, k) {
  const mati = mat.map((r, i) => [r, i])
  const hash = (r, i) => 100 * (r.reduce((a, b) => a + b)) + i
  mati.sort(([r1, i1], [r2, i2]) => hash(r1, i1) - hash(r2, i2))
  return mati.map(([r, i]) => i).slice(0, k)
};