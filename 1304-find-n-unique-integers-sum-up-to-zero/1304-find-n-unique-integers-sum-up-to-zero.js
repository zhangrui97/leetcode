/**
 * @param {number} n
 * @return {number[]}
 */
var sumZero = function(n) {
  const ans = []
  for (let i = 1; i <= ~~(n/2); i++) {
    ans.push(-i, i)
  }
  if (n%2) ans.push(0)
  return ans
};