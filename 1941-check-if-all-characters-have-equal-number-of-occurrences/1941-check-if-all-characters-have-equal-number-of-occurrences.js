/**
 * @param {string} s
 * @return {boolean}
 */
var areOccurrencesEqual = function(s) {
  const map = Array(26).fill(0)
  for (let ch of s) {
    map[ch.charCodeAt(0) - 97]++
  }
  const set = map.filter(n => n > 0)
  return set.every(n => n === set[0])
};