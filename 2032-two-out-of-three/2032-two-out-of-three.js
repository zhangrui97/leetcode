/**
 * @param {number[]} nums1
 * @param {number[]} nums2
 * @param {number[]} nums3
 * @return {number[]}
 */
var twoOutOfThree = function(nums1, nums2, nums3) {
  const map = new Map()
  nums1.forEach((n, i) => {
    if (nums1.indexOf(n) === i) map.set(n, 1)
  })
  nums2.forEach((n, i) => {
    if (nums2.indexOf(n) === i) {
      if (map.has(n))
        map.set(n, map.get(n) + 1)
      else map.set(n, 1)
    }
  })
  nums3.forEach((n, i) => {
    if (nums3.indexOf(n) === i) {
      if (map.has(n))
        map.set(n, map.get(n) + 1)
      else map.set(n, 1)
    }
  })
  const ans = []
  for (let [k, v] of map) {
    if (v > 1) ans.push(k)
  }
  return ans
};