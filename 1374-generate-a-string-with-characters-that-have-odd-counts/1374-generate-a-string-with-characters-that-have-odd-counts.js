/**
 * @param {number} n
 * @return {string}
 */
var generateTheString = function(n) {
  return n%2 ? Array(n).fill('x').join('') : 'x' + Array(n-1).fill('y').join('')
};