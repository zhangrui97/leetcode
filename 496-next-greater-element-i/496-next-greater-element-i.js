/**
 * @param {number[]} nums1
 * @param {number[]} nums2
 * @return {number[]}
 */
var nextGreaterElement = function(nums1, nums2) {
  const len = nums2.length
  return nums1.map(n => {
    for (let i = nums2.indexOf(n) + 1; i < len; i++) {
      if (nums2[i] > n) return nums2[i]
    }
    return -1
  })  
};