/**
 * @param {string} word
 * @return {number}
 */
var minTimeToType = function(word) {
  const len = word.length
  let ans = len
  let pos = 0
  const move = offset => offset > 13 ? 26 - offset : offset
  for (let ch of word) {
    const code = ch.charCodeAt(0) - 97
    ans += move((code + 26 - pos) % 26)
    pos = code
  }
  return ans
};