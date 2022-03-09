/**
 * @param {number[]} arr
 * @return {number[]}
 */
var sortByBits = function(arr) {
  const hamming = n => !n ? 0 : 1+hamming(n&(n-1))
  arr.sort((a, b) => (hamming(a)*100000+a - hamming(b)*100000-b))
  return arr
};