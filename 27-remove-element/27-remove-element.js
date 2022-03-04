/**
 * @param {number[]} nums
 * @param {number} val
 * @return {number}
 */
var removeElement = function(nums, val) {
  const len = nums.length
  if (len === 0) return 0
  if (len === 1) return nums[0] === val ? 0 : 1
  let fast = 0
  while (fast < len && nums[fast] != val) fast++
  let slow = fast
  while (fast < len) {
    while (fast < len && nums[fast] == val) fast++
    if (fast < len)
      nums[slow++] = nums[fast++]
  }
  return slow  
};