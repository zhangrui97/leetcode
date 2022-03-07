/**
 * @param {number[]} nums
 * @return {number[]}
 */
var sortArrayByParity = function(nums) {
  nums.sort((a, b) => {
    const aIsOdd = a % 2
    const bIsOdd = b % 2
    if (!aIsOdd && bIsOdd) return -1
    else if (aIsOdd && !bIsOdd) return 1
    else return a - b
  }) 
  return nums
};