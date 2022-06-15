/**
 * @param {number[]} nums
 * @return {void} Do not return anything, modify nums in-place instead.
 */
var nextPermutation = function(nums) {
  let n = nums.length
  let k = n - 2
  for (; k >= 0; k--) {
    if (nums[k] < nums[k+1]) {
      const v = nums[k]
      let l = n - 1
      for (; l > k+1; l--) {
        if (nums[l] > v) break;
      }
      nums[k] = nums[l]
      nums[l] = v
      for (let i = k+1, j = n-1; i < j; i++, j--) {
        [nums[i], nums[j]] = [nums[j], nums[i]]
      }
      return
    }
  }
  nums.reverse()
};