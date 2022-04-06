/**
 * @param {number[]} nums
 * @param {number} target
 * @return {number[]}
 */
var searchRange = function(nums, target) {
  let l = 0
  let r = nums.length - 1
  let mid = 0
  while (l <= r) {
    mid = ~~((l+r)/2)
    if (nums[mid] == target) break
    if (nums[mid] > target) {
      r = mid - 1
    } else {
      l = mid + 1
    }
  }
  if (l > r) return [-1, -1]
  let ll = l
  let lr = mid
  let m = l
  while (ll < lr) {
    m = ~~((ll+lr)/2)
    if (nums[m] < target) {
      ll = m + 1
    } else {
      lr = m
    }
  }
  let rl = mid
  let rr = r + 1
  while (rl < rr) {
    m = ~~((rl+rr)/2)
    if (nums[m] > target) {
      rr = m
    } else {
      rl = m + 1
    }
  }
  return [ll, rl-1]  
};