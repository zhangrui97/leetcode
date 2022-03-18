/**
 * @param {number} rowIndex
 * @return {number[]}
 */
var getRow = function(rowIndex) {
  const ans = new Array(rowIndex + 1).fill(1)
  let i = 1
  let f = rowIndex
  for (; i < ~~((rowIndex+1)/2); i++, f--) {
    const v = ans[i-1] * f / i
    ans[i] = v
    ans[rowIndex-i] = v
  }
  if (rowIndex > 0 && rowIndex % 2 === 0) ans[i] = ans[i-1] * f / i
  return ans
};