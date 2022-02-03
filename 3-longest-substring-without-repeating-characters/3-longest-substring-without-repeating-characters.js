/**
 * @param {string} s
 * @return {number}
 */
var lengthOfLongestSubstring = function(s) {
  const cache = new Map()
  let result = 0
  let l = 0
  let i = 0
  for (const c of s) {
    if (cache.has(c) && cache.get(c) >= l) {
      l = cache.get(c) + 1
    } else {
      const len = i + 1 - l
      if (len > result) result = len
    }
    cache.set(c, i)
    i += 1
  }
  return result
};