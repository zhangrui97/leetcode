/**
 * @param {number} x
 * @return {number}
 */
var reverse = function(x) {
  if (x < 0) return -reverse(-x)
  const result = +(('' + x).split('').reverse().join(''))
  return result >= 0x7FFFFFFF ? 0 : result
};