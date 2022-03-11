/**
 * @param {string} firstWord
 * @param {string} secondWord
 * @param {string} targetWord
 * @return {boolean}
 */
var isSumEqual = function(firstWord, secondWord, targetWord) {
  const convert = s => {
    return +([...s].map(ch => ch.charCodeAt(0) - 97).join(''))
  }
  return (convert(firstWord) + convert(secondWord)) === convert(targetWord)
};