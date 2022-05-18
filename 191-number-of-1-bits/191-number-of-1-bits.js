/**
 * @param {number} n - a positive integer
 * @return {number}
 */
var hammingWeight = function(n) {
  let i = 0;  
  for (; n != 0; i++) {
    n &= (n-1);
  }
  return i;
};