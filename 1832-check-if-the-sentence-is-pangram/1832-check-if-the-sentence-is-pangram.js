/**
 * @param {string} sentence
 * @return {boolean}
 */
var checkIfPangram = function(sentence) {
  const a = 'a'.charCodeAt(0)
  let flags = 0
  for (let ch of sentence) {
    flags |= 1 << (ch.charCodeAt(0) - a)
  }
  return (flags + 1) == 0x04000000
};