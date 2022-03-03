/**
 * @param {number[]} nums
 * @param {number} target
 * @return {number}
 */
var searchInsert = function(nums, target) {
  const search = (l, r) => {
    if (nums[l] >= target) return l
    if (nums[r] === target) return r
    if (nums[r] < target) return r + 1
    const mid = l + (~~((r-l) / 2))
    if (nums[mid] == target) return mid
    else if (nums[mid] < target) return search(mid + 1, r)
    else return search(l, mid - 1)
  }
  return search(0, nums.length - 1)
};