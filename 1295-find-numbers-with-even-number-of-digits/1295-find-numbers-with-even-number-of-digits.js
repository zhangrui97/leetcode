/**
 * @param {number[]} nums
 * @return {number}
 */
var findNumbers = function(nums) {
  const isEvenDigits = num => {
    if (num > 999) {
      return num <= 9999 || num == 100000
    } else {
      return num > 9 && num <= 99
    }
  }
  let ans = 0
  nums.forEach(n => ans += isEvenDigits(n))
  return ans
};