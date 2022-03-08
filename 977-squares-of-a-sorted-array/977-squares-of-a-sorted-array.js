/**
 * @param {number[]} nums
 * @return {number[]}
 */
var sortedSquares = function(nums) {
  nums.sort((a, b) => Math.abs(a) - Math.abs(b))
  return nums.map(n => n * n)
};