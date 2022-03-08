/**
 * @param {number[]} nums
 * @return {boolean}
 */
var canBeIncreasing = function(nums) {
  return nums.some((_, i) => {
    const arr = nums.slice()
    arr.splice(i, 1)
    return arr.every((n, i) => i === 0 || n > arr[i-1])
  })
};