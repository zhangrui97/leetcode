/**
 * @param {string} s
 * @return {number}
 */
var lengthOfLongestSubstring = function(s) {
  const cache = new Map()
  let result = 0
  let sub = new Set()
  let i = 0
  for (const c of s) {
    if (sub.has(c)) {
      const j = cache.get(c)
      sub = new Set(s.substr(j+1, i - j))
    } else {
      sub.add(c)
      if (sub.size > result) result = sub.size
    }
    cache.set(c, i)
    i += 1
  }
  return result
};