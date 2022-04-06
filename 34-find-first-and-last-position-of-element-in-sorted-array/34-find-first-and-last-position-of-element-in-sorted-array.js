/**
 * @param {number[]} nums
 * @param {number} target
 * @return {number[]}
 */
var searchRange = function(nums, target) {
  const fst = nums.indexOf(target)
  if (fst === -1) {
    return [-1, -1]
  } else {
    return [fst, nums.lastIndexOf(target)]
  }
};