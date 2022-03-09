/**
 * @param {string} jewels
 * @param {string} stones
 * @return {number}
 */
var numJewelsInStones = function(jewels, stones) {
  let ans = 0
  for (let ch of jewels) {
    ans += stones.split(ch).length - 1
  }
  return ans
};