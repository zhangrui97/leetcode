/**
 * @param {number[]} encoded
 * @param {number} first
 * @return {number[]}
 */
var decode = function(encoded, first) {
  const result = []
  result.push(first)
  let pre = first
  for (let n of encoded) {
    pre = pre ^ n
    result.push(pre)
  }
  return result
};