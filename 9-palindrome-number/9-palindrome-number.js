/**
 * @param {number} x
 * @return {boolean}
 */
var isPalindrome = function(x) {
  if(x < 0 || (x !== 0 && x % 10 == 0))
        return false
  const str = ''+x
  return str === [...str].reverse().join('')
};