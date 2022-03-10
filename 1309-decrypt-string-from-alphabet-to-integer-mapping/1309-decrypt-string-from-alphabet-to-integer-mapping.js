/**
 * @param {string} s
 * @return {string}
 */
var freqAlphabets = function(s) {
  const BASE = 'a'.charCodeAt(0) - 1
  let ans = []
  let i = 0
  let d = s.indexOf('#')
  while (i < s.length && d !== -1) {
    for (;i < d - 2; i++) {
      ans.push(+ s[i] + BASE)
    }
    ans.push(+ s[d-1] + BASE + 10 * s[d-2])
    i = d + 1
    d = s.indexOf('#', i)
  }
  for (;i < s.length; i++) {
    ans.push(+ s[i] + BASE)
  }
  return String.fromCharCode(...ans)
};