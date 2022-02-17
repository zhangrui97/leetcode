/**
 * @param {number[]} nums
 * @return {number[][]}
 */
var threeSum = function(nums) {
  if (nums.length < 3) return []
  nums.sort((a, b) => a - b)
  const len = nums.length
  if (nums[0] > 0 || nums[len-1] < 0) return []
  const twoSum = (i, target) => {
    const result = []
    let l = i
    let r = len - 1
    while (l < r) {
      const left = nums[l]
      const right = nums[r]
      const sum = left + right
      if (sum === target) {
        result.push([left, right])
        l ++
        r --
        while (nums[l] === left) l++
        while (nums[r] === right) r--
      } else if (sum < target) {
        l ++
        while (l < r && nums[l] === left) l++
      } else {// sum > target
        r --
        while (l < r && nums[r] === right) r--
      }
    }
    return result    
  }
  const ans = []
  for (let i = 0; i < len - 2;) {
    const a = nums[i]
    if (a > 0) break
    const res = twoSum(i+1, -a)
    if (res.length > 0) {
      const temp = res.map(arr => [a, ...arr])
      ans.push(...temp)
    }
    i++
    while (i < len - 2 && nums[i] === a) i++
  }
  return ans
};