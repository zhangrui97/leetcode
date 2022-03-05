/**
 * @param {number[]} nums
 * @return {number[]}
 */
var smallerNumbersThanCurrent = function(nums) {
  return nums.map(n => nums.filter(m => m < n).length)
};