/**
 * @param {string} s
 * @return {boolean}
 */
var areNumbersAscending = function(s) {
  const isDigit = ch => ch >= '0' && ch <= '9'
  let last = -1
  for (let i = 0; i < s.length;) {
    let digit = ''
    while (i < s.length && isDigit(s[i])) {
      digit += s[i]
      i ++
    }
    if (digit !== '') {
      const value = + digit
      if (value <= last) return false
      else {
        last = value
      }
    }
    i ++
  }
  return true
};