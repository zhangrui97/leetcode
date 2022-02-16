/**
 * @param {string} s
 * @return {number}
 */
var romanToInt = function(s) {
  const map = {I: 1, V: 5, X: 10, L: 50, C:100, D:500, M:1000}
  let sum = 0
  let i = 0
  for (; i < s.length - 1; i++) {
    if (map[s[i]] < map[s[i+1]]) {
      sum += map[s[i+1]] - map[s[i]]
      i++
    } else {
      sum += map[s[i]]
    }
  }
  if (i < s.length)
    return sum + map[s[i]]
  else return sum
};