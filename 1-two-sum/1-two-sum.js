/**
 * @param {number[]} nums
 * @param {number} target
 * @return {number[]}
 */
var twoSum = function(nums, target) {
    const rests = new Map()
    for (let i = 0; i < nums.length; i++) {
        if (rests.has(nums[i])) {
          return [rests.get(nums[i]), i]
        }
        rests.set(target - nums[i], i)
    }
};