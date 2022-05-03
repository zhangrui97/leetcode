/**
 * @param {number[]} nums1
 * @param {number[]} nums2
 * @return {number[]}
 */
var nextGreaterElement = function(nums1, nums2) {
  const map = {}
  nums1.forEach((v, i) => map[v] = i)
  const ans = Array(nums1.length).fill(-1)
  const stack = []
  for (const v of nums2) {
    while (stack.length > 0 && v > stack.at(-1)) {
      if (map[stack.at(-1)] != undefined) {
        ans[map[stack.at(-1)]] = v
      }
      stack.pop()
    }
    stack.push(v)
  }
  return ans
};