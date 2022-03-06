/**
 * @param {number[]} nums
 * @param {number} k
 * @return {number}
 */
var countPairs = function(nums, k) {
  const l = nums.length
  let ans = 0
  for (let i = 0; i < l - 1; i ++) {
    for (let j = i + 1; j < l; j++) {
      if ((i*j % k) === 0 && nums[i] === nums[j]) {
        ans ++
      }
    }
  }
  return ans
};