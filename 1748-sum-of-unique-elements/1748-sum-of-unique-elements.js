/**
 * @param {number[]} nums
 * @return {number}
 */
var sumOfUnique = function(nums) {
  const set = new Set(nums)
  nums.forEach((n, i) => nums.indexOf(n) != i ? set.delete(n) : 0)
  if (set.size === 0) return 0
  return [...set].reduce((a, b) => a + b)
};