/**
 * @param {number[]} nums
 * @return {number[]}
 */
var nextGreaterElements = function(nums) {
  const stack = []
  const ans = Array(nums.length).fill(-1)
  nums.forEach((v, i) => {
    while (stack.length > 0 && v > nums[stack.at(-1)]) {
      ans[stack.pop()] = v
    }
    stack.push(i)
  })
  nums.forEach((v, i) => {
    while (stack.length > 0 && v > nums[stack.at(-1)]) {
      ans[stack.pop()] = v
    }
  })
  return ans
};