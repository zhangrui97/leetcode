/**
 * @param {number[]} nums
 * @return {number}
 */
var findNumbers = function(nums) {
  let ans = 0
  nums.forEach(n => ans += (('' + n).length % 2 === 0))
  return ans
};