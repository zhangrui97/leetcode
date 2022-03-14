/**
 * @param {string} sentence
 * @param {string} searchWord
 * @return {number}
 */
var isPrefixOfWord = function(sentence, searchWord) {
  let ans = -1
  const a = sentence.split(' ')
  for (let i = 0; i < a.length; i++) {
    if (a[i].startsWith(searchWord)) return i+1
  }
  return ans
};