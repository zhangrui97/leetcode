/**
 * @param {string} s
 * @param {number} k
 * @return {number}
 */
var getLucky = function(s, k) {
  let result = ''
  for (let ch of s) {
    result += ch.charCodeAt(0) - 96
  }
  for (let i = 0; i < k; i ++) {
    let value = 0
    for (let ch of result) {
      value += +ch
    }
    result = '' + value
  }
  return +result
};