/**
 * @param {string} s
 * @return {number}
 */
var countGoodSubstrings = function(s) {
  if (s.length < 3) return 0
  let count = 0
  let diff = s[0] !== s[1]
  for (let i = 2; i < s.length; i++) {
    if (diff && s[i] != s[i-2]) {
      if (s[i] !== s[i-1])
        count++
      else diff = false
    } else {
      diff = s[i] !== s[i-1]
    }
  }
  return count
};