/**
 * @param {number} x
 * @return {boolean}
 */
var isPalindrome = function(x) {
  if(x < 0 || (x !== 0 && x % 10 == 0))
        return false
  const isPal = s => 
    s.length < 2 || (s[0] === s.at(-1) && isPal(s.slice(1, -1)))
  return isPal(''+x)
};