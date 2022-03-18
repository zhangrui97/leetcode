/**
 * @param {number} rowIndex
 * @return {number[]}
 */
var getRow = function(rowIndex) {
  switch (rowIndex) {
    case 0: return [1]
    case 1: return [1, 1]
    case 2: return [1, 2, 1]
    case 3: return [1, 3, 3, 1]
    case 4: return [1, 4, 6, 4, 1]
    default:
      const ans = new Array(rowIndex + 1)
      ans[0] = 1
      ans[rowIndex] = 1
      let i = 1
      let f = rowIndex
      for (; i < ~~((rowIndex+1)/2); i++, f--) {
        const v = ans[i-1] * f / i
        ans[i] = v
        ans[rowIndex-i] = v
      }
      if (rowIndex % 2 === 0) ans[i] = ans[i-1] * f / i
      return ans
  }
};