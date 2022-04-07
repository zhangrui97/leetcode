/**
 * @param {number[]} nums
 * @param {number} k
 * @return {number}
 */
var subarraySum = function(nums, k) {
  const map = new Map()
  map.set(0, 1)
  let sum = 0
  let ans = 0
  nums.forEach( v => {
    sum += v
    if (map.has(sum-k)) {
      ans += map.get(sum-k)
    }
    map.set(sum, map.has(sum) ? map.get(sum)+1 : 1)
  })
  return ans
};