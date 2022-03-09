/**
 * @param {string} s
 * @param {character} c
 * @return {number[]}
 */
var shortestToChar = function(s, c) {
  const ans = Array(s.length)
  let pre = s.indexOf(c)
  let last = s.lastIndexOf(c)
  for (let i = 0; i <= pre; i++) {
    ans[i] = pre - i
  }
  next = pre + 1
  while (next < last) {
    while (s[next] != c) next++
    for (let i = pre + 1; i <= next; i++) {
      ans[i] = Math.min(i - pre, next - i)
    }
    pre = next
    next++
  }
  for (let i = s.length - 1; i >= last; i--) {
    ans[i] = i - last
  }
  return ans
};