/**
 * @param {number[]} nums
 * @return {number}
 */
var lengthOfLIS = function(nums) {
  const heap = []
  for (const n of nums) {
    let [l, r] = [0, heap.length]
    while (l < r) {
      const m = ~~((l + r) / 2)
      if (heap[m] < n) l = m + 1
      else r = m
    }
    if (l >= heap.length) heap.push(n)
    else heap[l] = n
  }
  return heap.length
};