/**
 * @param {number[]} nums
 * @return {number}
 */
var numIdenticalPairs = function(nums) {
  const l = nums.length
  let count = 0
  for (let i = 0; i < l - 1; i++) {
    for (let j = i + 1; j < l; j++) {
      if (nums[i] === nums[j]) count++
    }
  }
  return count
};