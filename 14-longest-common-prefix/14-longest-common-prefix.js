/**
 * @param {string[]} strs
 * @return {string}
 */
var longestCommonPrefix = function(strs) {
  const len = Math.min.apply(null, strs.map(s => s.length))
  let result = ''
  for (let i = 0; i < len; i++) {
    const ch = strs[0][i]
    if (!strs.every(s => s[i] === ch)) break
    result += ch
  }
  return result
};