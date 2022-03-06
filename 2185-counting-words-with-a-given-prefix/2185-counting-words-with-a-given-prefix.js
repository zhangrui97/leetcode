/**
 * @param {string[]} words
 * @param {string} pref
 * @return {number}
 */
var prefixCount = function(words, pref) {
  let ans = 0
  words.forEach(word => ans += word.startsWith(pref))
  return ans
};