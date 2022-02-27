const map = {'(':')', '[':']', '{':'}'}
/**
 * @param {string} s
 * @return {boolean}
 */
var isValid = function(s) {
  if (s.length === 1) return false
  let open = []
  for (let ch of s) {
    if (ch in map) {
      open.push(map[ch])
    } else {
      if (open.pop() != ch)
        return false
    }
  }
  return open.length === 0
};