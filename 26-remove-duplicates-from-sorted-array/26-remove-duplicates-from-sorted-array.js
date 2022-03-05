/**
 * @param {number[]} nums
 * @return {number}
 */
var removeDuplicates = function(nums) {
  if (nums.length === 1) return 1
  let fast = 1
  let slow = 0
  while (fast < nums.length) {
    if (nums[fast] === nums[slow]) fast++
    else {
      slow += 1
      if (fast !== slow)
        nums[slow] = nums[fast]
      fast += 1
    }
  }
  return slow + 1
};