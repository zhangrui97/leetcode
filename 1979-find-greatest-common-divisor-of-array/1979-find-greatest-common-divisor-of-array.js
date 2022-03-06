/**
 * @param {number[]} nums
 * @return {number}
 */
var findGCD = function(nums) {
  const gcd = (a, b) => (a > b) ? gcd(b, a) : (a === 0) ? b : gcd(b % a, a)
  return gcd(Math.min(...nums), Math.max(...nums))
};