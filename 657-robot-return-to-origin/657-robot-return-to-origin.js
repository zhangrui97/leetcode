/**
 * @param {string} moves
 * @return {boolean}
 */
var judgeCircle = function(moves) {
  let o = {U: 0, D: 0, L: 0, R: 0}
  for (let ch of moves) o[ch]++
  return o.U === o.D && o.L === o.R
};