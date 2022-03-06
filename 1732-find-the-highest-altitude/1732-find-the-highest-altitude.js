/**
 * @param {number[]} gain
 * @return {number}
 */
var largestAltitude = function(gain) {
  const len = gain.length
  const alts = Array(len + 1)
  alts[0] = 0
  for (let i = 0; i < len; i++) {
    alts[i+1] = alts[i] + gain[i]
  }
  return Math.max(...alts)
};