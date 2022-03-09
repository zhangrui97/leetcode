/**
 * @param {number[]} nums
 * @return {number[]}
 */
var sortArrayByParityII = function(nums) {
  const odd = nums.filter(n => n%2)
  const even = nums.filter(n => n%2 === 0)
  console.log(odd, even)
  const ans = Array(nums.length)
  for (let i = 0; i < odd.length; i++) {
    ans[2 * i] = even[i]
    ans[2 * i + 1] = odd[i]
  }
  return ans
};