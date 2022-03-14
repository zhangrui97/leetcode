/**
 * @param {string[]} words
 * @return {string[]}
 */
var stringMatching = function(words) {
  return words.filter((word, i) => words.some((w, j) => i !== j && w.indexOf(word) !== -1))  
};