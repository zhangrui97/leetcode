/**
 * @param {number[]} nums
 * @param {number} original
 * @return {number}
 */
var findFinalValue = function(nums, original) {
  if (!original) return 0
  while (nums.find(el => el === original)) {
    original += original
  }
  return original
};