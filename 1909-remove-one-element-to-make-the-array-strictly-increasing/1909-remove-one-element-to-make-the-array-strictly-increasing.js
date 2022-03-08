/**
 * @param {number[]} nums
 * @return {boolean}
 */
var canBeIncreasing = function(nums) {
  let once = 0
  for (let i = 1; i < nums.length; i ++) {
    if (nums[i] <= nums[i-1]) {
      if (i === 1 || i === nums.length - 1) once += 1
      else if (nums[i] > nums[i-2] || nums[i+1] > nums[i-1]) once += 1
      else return false
    }
    if (once > 1) return false
  }
  return true
};