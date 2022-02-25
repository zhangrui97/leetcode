/**
 * @param {number[]} nums
 * @param {number} target
 * @return {number}
 */
var threeSumClosest = function(nums, target) {
  if (nums.length <= 3) return nums.reduce((a, b) => a + b)
  nums.sort((a, b) => a - b)
  if (nums[0] >= (target/3 + 1)) return nums[0] + nums[1] + nums[2]
  if (nums[nums.length-1] <= target/3) return nums[nums.length-3] + nums[nums.length-2] + nums[nums.length-1]
  var min = 2000
  for (let i = 0; i < nums.length-2; i++) {
    const a = nums[i]
    const rest = target - a
    var [l, r] = [i+1, nums.length -1]
    while (l < r) {
      const sum = nums[l] + nums[r]
      if (sum < rest) {
        l++
      } else if (sum === rest) {
        return target
      } else {
        r--
      }
      if (Math.abs(sum - rest) < Math.abs(min - target)) {
        min = a + sum
      }
    }
  }
  return min
};