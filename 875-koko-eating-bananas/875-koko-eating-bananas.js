/**
 * @param {number[]} piles
 * @param {number} h
 * @return {number}
 */
var minEatingSpeed = function(piles, h) {
  const eat = speed => piles.map(v => Math.ceil(v/speed)).reduce((a, b) => a+b)
  let min = 1
  let max = Math.max(...piles) + 1
  while (min < max) {
    const mid = min + ~~((max - min) / 2)
    if (eat(mid) > h) min = mid + 1
    else max = mid
  }
  return min
};