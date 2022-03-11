/**
 * @param {string} word
 * @param {character} ch
 * @return {string}
 */
var reversePrefix = function(word, ch) {
  let stack = []
  let i = 0
  for (; i < word.length; i++) {
    stack.push(word[i])
    if (word[i] === ch) break
  }
  return (i === word.length)
    ? word
    : stack.reverse().join('') + word.substring(i+1)
};