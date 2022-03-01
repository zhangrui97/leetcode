/**
 * @param {number} n
 * @return {string[]}
 */
var generateParenthesis = function(n) {
  const result = []
  const backTrack = (s, l, r) => {
    if (l === 0 && r === 0) {
      result.push(s)
    }
    if (l > 0) backTrack(s + '(', l-1, r)
    if (l < r) backTrack(s + ')', l, r-1)
  }
  backTrack('', n, n);
  return result;
};