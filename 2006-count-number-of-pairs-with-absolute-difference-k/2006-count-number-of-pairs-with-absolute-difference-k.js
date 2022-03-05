/**
 * @param {number[]} nums
 * @param {number} k
 * @return {number}
 */
var countKDifference = function(nums, k) {
  const l = nums.length
  let count = 0
  for (let i = 0; i < l - 1; i++) {
    for (let j = i + 1; j < l; j++) {
      const delta = nums[i] - nums[j]
      if (delta === k || delta === -k) count++
    }
  }
  return count
};