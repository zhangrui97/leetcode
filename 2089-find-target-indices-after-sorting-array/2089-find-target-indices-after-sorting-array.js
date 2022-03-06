/**
 * @param {number[]} nums
 * @param {number} target
 * @return {number[]}
 */
var targetIndices = function(nums, target) {
  const ans = []
  nums.sort((a, b) => a - b).forEach((n, i) => (n === target) ? ans.push(i) : 0)
  return ans
};