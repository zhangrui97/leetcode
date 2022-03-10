/**
 * @param {string} s
 * @return {string}
 */
var removeOuterParentheses = function(s) {
  let ans = ''
  let depth = 0
  for (let ch of s) {
    if (ch === '(') {
      if (depth > 0) {
        ans += ch
      }
      depth++
    } else {
      depth--
      if (depth > 0) {
        ans += ch
      }
    }
  }    
  return ans
};