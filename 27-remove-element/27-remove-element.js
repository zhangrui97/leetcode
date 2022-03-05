/**
 * @param {number[]} nums
 * @param {number} val
 * @return {number}
 */
var removeElement = function(nums, val) {
  const len = nums.length
  if (len === 0) return 0
  let fast = 0
  while (fast < len && nums[fast] != val) fast++
  let slow = fast
  while (fast < len) {
    if (nums[fast] == val) fast++
    else
      nums[slow++] = nums[fast++]
  }
  return slow  
};