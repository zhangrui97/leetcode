/**
 * @param {string} s1
 * @param {string} s2
 * @return {boolean}
 */
var checkInclusion = function(s1, s2) {
  const counter = {}  
  for (const ch of s1) {
    counter[ch] = (counter[ch] || 0) + 1
  }
  const needs = Object.keys(counter).length
  let l = 0
  let r = 0
  let valid = 0
  const ctx = {}
  while (r < s2.length) {
    const chr = s2[r]
    r++
    if (!counter[chr]) continue
    ctx[chr] = (ctx[chr] || 0) + 1
    if (ctx[chr] !== counter[chr]) continue
    valid++
    while (valid === needs) {
      if (r - l === s1.length) return true
      const chl = s2[l]
      if (counter[chl]) {
        ctx[chl] -= 1
        if (ctx[chl] < counter[chl]) {
          valid--
        }
      }
      l++
    }
  }
  return false
};