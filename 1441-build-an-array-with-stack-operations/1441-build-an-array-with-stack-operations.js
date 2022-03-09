/**
 * @param {number[]} target
 * @param {number} n
 * @return {string[]}
 */
var buildArray = function(target, n) {
  const ans = []
  const build = (start, end) => {
    while (start < end) {
      ans.push("Push")
      ans.push("Pop")
      start ++
    }
    ans.push("Push")
  }
  let pre = 1
  target.forEach(n => {
    build(pre, n)
    pre = n + 1
  })
  return ans
};