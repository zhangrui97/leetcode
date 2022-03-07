/**
 * @param {number[]} nums
 * @return {number}
 */
var arrayPairSum = function(nums) {
  nums.sort((a, b) => a - b)
  let ans = 0
  nums.forEach((n, i) => (i%2 === 0) ? ans += n : 0)
  return ans
};