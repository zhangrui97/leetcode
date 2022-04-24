let cnt = 0
/**
 * @param {number[]} nums
 * @param {number} k
 * @return {number}
 */
var subarraySum = function(nums, k) {

  let counter = 0
  const map = new Map()
  map.set(0, 1)
  let sum = 0
  for (let i = 0; i < nums.length; i++) {
    sum += nums[i]
    if (map.has(sum-k)) {
      counter += map.get(sum-k)
    }
    map.set(sum, map.has(sum) ? map.get(sum)+1 : 1)
  }
  return counter
}