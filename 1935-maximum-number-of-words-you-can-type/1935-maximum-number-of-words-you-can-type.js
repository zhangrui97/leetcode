/**
 * @param {string} text
 * @param {string} brokenLetters
 * @return {number}
 */
var canBeTypedWords = function(text, brokenLetters) {
  return text.split(' ').filter(s => {
    for (let ch of s) {
      if (brokenLetters.includes(ch))
        return false
    }
    return true
  }).length 
};