/**
 * @param {string[]} strs
 * @return {number}
 */
var minDeletionSize = function(strs) {
  if (strs.length <= 1) return 0
  const len = strs[0].length
  let count = 0;
  for (let j = 0; j < len; j++) {
    for (let i = 1; i < strs.length; i++) {
      if (strs[i][j] < strs[i-1][j]) {
        count++
        break
      }
    }
  }
  return count
};