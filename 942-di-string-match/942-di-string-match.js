/**
 * @param {string} s
 * @return {number[]}
 */
var diStringMatch = function(s) {
  const ans = [0]
  let min = 0
  let max = 0
  for (let ch of s) {
    if (ch === 'I') {
      ans.push(++max)
    } else {
      ans.push(--min)
    }
  }
  return ans.map(e => e - min)
};