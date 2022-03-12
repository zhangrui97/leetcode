/**
 * @param {string} s
 * @return {number}
 */
var removePalindromeSub = function(s) {
  if (s.length === 0) return 0
  let l = 0
  let r = s.length -1
  while (l < r && s[l] === s[r]) {
    l++
    r--
  }
  if (l < r) return 2
  return 1
};