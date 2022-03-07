/**
 * @param {number[]} heights
 * @return {number}
 */
var heightChecker = function(heights) {
  const sorted = heights.slice()
  sorted.sort((a, b) => a - b)
  let ans = 0
  for (let i = 0; i < heights.length; i++) {
    ans += heights[i] != sorted[i]
  }
  return ans
};