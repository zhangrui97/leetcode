/**
 * @param {number[]} nums
 * @param {number} target
 * @return {number[]}
 */
var twoSum = function(nums, target) {
  for (i = 0; i < nums.length; i++) {
    for (j = 0; j < i; j++) {
      if (nums[j] + nums[i] === target) {
        return [i, j]
      }
    }
  }
};