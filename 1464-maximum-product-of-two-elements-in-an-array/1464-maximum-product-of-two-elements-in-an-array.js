/**
 * @param {number[]} nums
 * @return {number}
 */
var maxProduct = function(nums) {
  const fst = Math.max(...nums)
  nums.splice(nums.indexOf(fst), 1)
  const snd = Math.max(...nums)
  return (fst - 1) * (snd - 1)
};