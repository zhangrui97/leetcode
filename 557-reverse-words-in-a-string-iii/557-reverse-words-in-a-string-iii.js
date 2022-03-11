/**
 * @param {string} s
 * @return {string}
 */
var reverseWords = function(s) {
  return s.split(' ').map(word => [...word].reverse().join('')).join(' ')
};