/**
 * @param {number[]} nums
 * @return {number}
 */
var subsetXORSum = function(nums) {
  if (nums.length === 0) return 0
  const xorTotal = a => {
    if (a.length === 0) return 0
    if (a.length === 1) return a[0]
    return a.reduce((e1, e2) => e1 ^ e2)
  }
  let ans = 0
  const addSub = a => {
    if (a.length === 0) return []
    if (a.length === 1) {
      ans += a[0]
      return [a, []]
    }
    const [h, ...t] = a
    const sub1 = addSub(t)
    const sub2 = sub1.map(a => [h, ...a])
    sub2.forEach(a => ans += xorTotal(a))
    return [...sub1, ...sub2]
  }
  addSub(nums)
  return ans
};