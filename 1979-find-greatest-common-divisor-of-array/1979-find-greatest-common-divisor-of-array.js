/**
 * @param {number[]} nums
 * @return {number}
 */
var findGCD = function(nums) {
  const gcd = (a, b) => a ? gcd(b % a, a) : b
  return gcd(Math.min(...nums), Math.max(...nums))
};