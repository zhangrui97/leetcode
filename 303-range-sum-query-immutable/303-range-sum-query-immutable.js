/**
 * @param {number[]} nums
 */
var NumArray = function(nums) {
  this.sums = Array(nums.length)
  let sum = 0
  nums.forEach((v, i) => {
    sum += v
    this.sums[i] = sum
  })
};

/** 
 * @param {number} left 
 * @param {number} right
 * @return {number}
 */
NumArray.prototype.sumRange = function(left, right) {
  return this.sums[right] - (left === 0 ? 0 : this.sums[left-1])
}

/** 
 * Your NumArray object will be instantiated and called as such:
 * var obj = new NumArray(nums)
 * var param_1 = obj.sumRange(left,right)
 */