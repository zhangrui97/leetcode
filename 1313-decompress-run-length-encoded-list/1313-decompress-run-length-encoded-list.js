/**
 * @param {number[]} nums
 * @return {number[]}
 */
var decompressRLElist = function(nums) {
  const result = []
  for (let i = 0; i < nums.length / 2; i++) {
    const len = nums[i*2]
    const val = nums[i*2+1]
    for (let j = 0; j < nums[i*2]; j++) {
      result.push(val)
    }
  }
  return result
};