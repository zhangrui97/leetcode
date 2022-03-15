/**
 * @param {number} n
 * @return {number[]}
 */
var countBits = function(n) {
  const countOne = x => {
    let count = 0
    while (x) {
      count++
      x &= x - 1
    }
    return count
  } 
  return new Array(n+1).fill(0).map((_, i) => countOne(i))
};