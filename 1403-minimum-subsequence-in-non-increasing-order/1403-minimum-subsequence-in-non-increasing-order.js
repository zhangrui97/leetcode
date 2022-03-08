/**
 * @param {number[]} nums
 * @return {number[]}
 */
var minSubsequence = function(nums) {
  if (nums.length < 2) return nums
  const sum = nums.reduce((a, b) => a+b)
  nums.sort((a, b) => b - a)
  const ans = []
  let rest = sum / 2
  for (let i = 0; i < nums.length && rest >= 0; i++) {
    const n = nums[i]
    ans.push(n)
    rest -= n
  }
  return ans
};