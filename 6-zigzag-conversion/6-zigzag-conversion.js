/**
 * @param {string} s
 * @param {number} numRows
 * @return {string}
 */
var convert = function(s, numRows) {
  let len = s.length
  if (numRows <= 1) return s
  let step = 2 * numRows - 2
  let result = ''
  for (let i = 0; i < numRows; i++) {
    let index = i;
    let index2 = i%(numRows-1) ? step - i : len;
    while (index < len || index2 < len) {
      if (index < len) {
        result += s[index]
        index += step
      }
      if (index2 < len) {
        result += s[index2]
        index2 += step
      }
    }
  }
  return result;
};