/**
 * @param {number[]} nums
 * @return {number[]}
 */
var runningSum = function(nums) {
  const sum = (pre, res) => {
    if (res.length === 0) return []
    const [h, ...t] = res
    const nxt = h + pre
    return [nxt, ...sum(nxt, t)]
  }
  return sum(0, nums)
};