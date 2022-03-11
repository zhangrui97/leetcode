/**
 * @param {number} n
 * @return {string}
 */
var generateTheString = function(n) {
  const dict = "abcdefghi"
  const generator = (m, i) => {
    if (m === 1) return [dict[i]]
    if ((m & 3) === 3) {
      return new Array(m >> 1).fill(dict[i]).concat(generator((m >> 1) + 1, i+1))
    } else if ((m & 1) === 1) {
      return new Array((m >> 1) + 1).fill(dict[i]).concat(generator(m >> 1, i+1))
    } else if ((m & 3) === 0) {
      return new Array((m >> 1) + 1).fill(dict[i]).concat(new Array((m >> 1) - 1).fill(dict[i+1]))
    } else { // if (m & 1 === 0)
      return new Array(m >> 1).fill(dict[i]).concat(new Array(m >> 1).fill(dict[i+1]))
    }
  }
  return generator(n, 0).join('')
};