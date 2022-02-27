const map = new Map([[')','('], [']','['], ['}','{']])
const set = new Set(map.values())
/**
 * @param {string} s
 * @return {boolean}
 */
var isValid = function(s) {
  if (s.length === 1) return false
  let open = []
  let isMatch = ch => open.length > 0 && map.has(ch) && open[0] === map.get(ch)
  for (let ch of s) {
    if (set.has(ch)) {
      open = ch + open
    } else {
      if (!isMatch(ch)) return false 
      open = open.substring(1)
    }
  }
  return open == []
};