/**
 * @param {string} text
 * @return {number}
 */
var maxNumberOfBalloons = function(text) {
  const map = { b: 0, a: 0, l: 0, o: 0, n: 0}
  for (let ch of text) {
    if (ch in map) map[ch]++
  }
  return Math.min(map.b, map.a, ~~(map.l / 2), ~~(map.o / 2), map.n)
};