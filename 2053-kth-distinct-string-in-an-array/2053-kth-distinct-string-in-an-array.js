/**
 * @param {string[]} arr
 * @param {number} k
 * @return {string}
 */
var kthDistinct = function(arr, k) {
  let count = 0
  for (let s of arr) {
    count += (arr.indexOf(s) == arr.lastIndexOf(s))
    if (count === k) return s
  }
  return ''
};