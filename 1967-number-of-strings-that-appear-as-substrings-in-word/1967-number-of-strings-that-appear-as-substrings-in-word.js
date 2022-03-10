/**
 * @param {string[]} patterns
 * @param {string} word
 * @return {number}
 */
var numOfStrings = function(patterns, word) {
  return patterns.filter(pat => word.includes(pat)).length    
};