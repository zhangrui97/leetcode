/**
 * @param {number[]} position
 * @return {number}
 */
var minCostToMoveChips = function(position) {
  let odd = 0, even = 0
  position.forEach(n => n % 2 ? odd++ : even++)
  return Math.min(odd, even)
};