/**
 * @param {number[]} height
 * @return {number}
 */
var maxArea = function(height) {
    let left = 0
    let right = height.length -1
    let result = 0
    while (left < right) {
      const l = height[left]
      const r = height[right]
      if (l < r) {
        result = Math.max(result, (right-left)*l)
        left++
      } else if (r < l) {
        result = Math.max(result, (right-left)*r)
        right--
      } else {
        result = Math.max(result, (right-left)*l)
        left++
        right--
      }
    }
    return result
};