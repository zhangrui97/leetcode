/**
 * @param {number[]} nums
 * @param {number} k
 * @return {boolean}
 */
var canPartitionKSubsets = function(nums, k) {
  const n = nums.length
  if (k > n) { return false }
  const sum = nums.reduce((a, b) => a+b, 0)
  if (sum % k != 0) { return false }
  const mean = sum / k
  nums.sort((a, b) => a - b);
  if (nums.at(-1) > mean) { return false }
  let used = 0
  const map = {}
  const backtrack = (bucket, start, target) => {
    if (bucket === k) return true
    if (target === 0) {
      const res = backtrack(bucket+1, 0, mean)
      map[used] = res
      return res
    }
    if (map[used] !== undefined) { return map[used] }
    let pre = 0
    for (let i = start; i < n; i++) {
      if (used & (1 << i)) continue
      const v = nums[i]
      if (v > target) return false
      if (v === pre) continue
      used |= (1 << i)
      pre = v
      const res = backtrack(bucket, i + 1, target - v)
      map[used] = res
      used &= ~(1 << i)
      if (res) return true
    }
    return false
  }
  return backtrack(0, 0, mean)
};