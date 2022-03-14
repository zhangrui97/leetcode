/**
 * @param {string} s
 * @return {number}
 */
var maxPower = function(s) {
  let ans = 1
  for (let i = 0; i < s.length;) {
    let j = i+1
    while (j < s.length && s[j] === s[i]) {
      j ++
      ans = Math.max(ans, j - i)
    }
    i = j
  }
  return ans
};