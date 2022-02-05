/**
 * @param {number} x
 * @return {boolean}
 */
var isPalindrome = function(x) {
    const helper = s => s.length < 2 || (s[0] === s.at(-1) && isPalindrome(s.slice(1, -1)))
    return helper(''+x)
};