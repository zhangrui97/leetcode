/**
 * @param {number[]} nums
 * @return {number}
 */
var sumOfUnique = function(nums) {
  let ans = 0
  nums.forEach(n => ans += nums.indexOf(n) == nums.lastIndexOf(n) ? n : 0)
  return ans
};