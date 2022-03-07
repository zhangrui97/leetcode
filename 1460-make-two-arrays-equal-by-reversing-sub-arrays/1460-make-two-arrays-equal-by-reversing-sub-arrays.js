/**
 * @param {number[]} target
 * @param {number[]} arr
 * @return {boolean}
 */
var canBeEqual = function(target, arr) {
  if (target.length !== arr.length) return false
  const comp = (a, b) => a - b
  target.sort(comp)
  arr.sort(comp)
  for (let i = 0; i < target.length; i++) {
    if (arr[i] !== target[i]) return false
  }
  return true
};