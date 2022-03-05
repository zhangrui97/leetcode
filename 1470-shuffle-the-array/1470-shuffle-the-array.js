/**
 * @param {number[]} nums
 * @param {number} n
 * @return {number[]}
 */
var shuffle = function(nums, n) {
  const a1 = nums.slice(0, n)
  const a2 = nums.slice(n)
  const result = []
  for (let i = 0; i < n; i++) {
    result.push(a1[i])
    result.push(a2[i])
  }
  return result
};