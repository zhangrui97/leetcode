/**
 * @param {string} s
 * @return {number}
 */
var maxDepth = function(s) {
  let ans = 0
  let depth = 0
  for (let ch of s) {
    if (ch === `(`) depth++
    else if (ch === `)`) depth--
    ans = Math.max(ans, depth)
  }
  return ans
};