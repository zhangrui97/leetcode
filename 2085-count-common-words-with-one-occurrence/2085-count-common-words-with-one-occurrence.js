/**
 * @param {string[]} words1
 * @param {string[]} words2
 * @return {number}
 */
var countWords = function(words1, words2) {
  const once = words1.filter(n => words1.indexOf(n) === words1.lastIndexOf(n))
  let count = 0
  words2.forEach(n => count += once.includes(n) && words2.indexOf(n) === words2.lastIndexOf(n))
  return count
};