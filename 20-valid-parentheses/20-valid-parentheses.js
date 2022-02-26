/**
 * @param {string} s
 * @return {boolean}
 */
var isValid = function(s) {
  if (s.length === 1) return false
  let open = []
  let isMatch = ch => {
    if (open.length <= 0) return false
    switch (ch) {
      case ')': return open[0] === '('
      case ']': return open[0] === '['
      case '}': return open[0] === '{'
    }
  }
  for (let ch of s) {
    if (ch === '(' || ch === '[' || ch === '{') {
      open = ch + open
    } else {
      if (!isMatch(ch)) return false 
      open = open.substring(1)
    }
  }
  return open == []
};