/**
 * @param {number[]} nums
 * @param {number} k
 * @return {number}
 */
var subarraySum = function(nums, k) {
  const cache = {}
  cache[0] = 1
  let sum = 0
  let count = 0
  for (const n of nums) {
    sum += n
    if (cache[sum-k]) {
      count += cache[sum-k]
    }
    if (cache[sum]) {
      cache[sum]++
    } else {
      cache[sum]  = 1
    }
  }
  return count
}