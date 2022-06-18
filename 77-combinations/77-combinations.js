/**
 * @param {number} n
 * @param {number} k
 * @return {number[][]}
 */
var combine = function(n, k) {
  const result = []
  const dfs = (path, i, l) => {
    if (l === k) {
      const res = []
      for (let j = 1, p = path>>1; p > 0; j++, p = p>>1) {
        if (p & 1) res.push(j)
      }
      result.push(res)
    } else {
      for (let j = i; j <= n; j++) {
        dfs(path | (1 << j), j+1, l+1)
      }
    }
  }
  dfs(0, 1, 0)
  return result
};