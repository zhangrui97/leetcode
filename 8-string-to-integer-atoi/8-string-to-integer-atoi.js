/**
 * @param {string} s
 * @return {number}
 */
var myAtoi = function(s) {
  const MIN = -0x80000000
  const MAX = 0x80000000 - 1
  const regex = /^\s*([+-]?\d+)[\s\.\d\-\+a-zA-Z]*$/
  if (regex.test(s)) {
    const result = +(s.replace(regex, '$1'))
    return Math.min(Math.max(result, MIN), MAX)
  }
  else return 0
};