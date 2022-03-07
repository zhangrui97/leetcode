/**
 * @param {number[]} nums
 * @return {number}
 */
var repeatedNTimes = function(nums) {
  for (let i = 0; i <= nums.length / 2; i++) {
    const n = nums[i]
    if (nums.lastIndexOf(n) != i) return n
  }
  return 0;
};