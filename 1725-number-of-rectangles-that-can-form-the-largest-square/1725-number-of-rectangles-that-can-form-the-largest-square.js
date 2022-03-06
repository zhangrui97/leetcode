/**
 * @param {number[][]} rectangles
 * @return {number}
 */
var countGoodRectangles = function(rectangles) {
  const squares = rectangles.map(a => Math.min(...a))
  const max = Math.max(...squares)
  let ans = 0
  squares.forEach(l => (l === max) ? ans++ : 0)
  return ans
};