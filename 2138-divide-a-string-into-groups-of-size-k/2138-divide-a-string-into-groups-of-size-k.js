/**
 * @param {string} s
 * @param {number} k
 * @param {character} fill
 * @return {string[]}
 */
var divideString = function(s, k, fill) {
  const ans = []
  let i = 0
  for (; i < s.length - k; i += k) {
    ans.push(s.substring(i, i + k))
  }
  if (i !== s.length) {
    ans.push(s.substring(i) + (new Array(i + k - s.length)).fill(fill).join(''))
  }
  return ans
};