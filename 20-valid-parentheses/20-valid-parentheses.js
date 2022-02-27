const map = new Map([['(',')'], ['[',']'], ['{','}']])
/**
 * @param {string} s
 * @return {boolean}
 */
var isValid = function(s) {
  if (s.length === 1) return false
  let open = []
  let isMatch = ch => open.length > 0 && open.at(-1) === ch
  for (let ch of s) {
    if (map.has(ch)) {
      open.push(map.get(ch))
    } else {
      if (!isMatch(ch)) return false 
      open.pop()
    }
  }
  return open.length === 0
};