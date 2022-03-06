/**
 * @param {string[]} words
 * @return {string}
 */
var firstPalindrome = function(words) {
  const isPal = s => s.length < 2 || (s[0] === s.at(-1) && isPal(s.slice(1, -1)))
  for (let word of words) {
    if (isPal(word)) return word
  }
  return ""
};