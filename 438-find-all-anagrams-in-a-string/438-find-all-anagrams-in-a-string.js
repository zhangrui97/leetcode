/**
 * @param {string} s
 * @param {string} p
 * @return {number[]}
 */
var findAnagrams = function(s, p) {
  if (p.length > s.length) return []
  const counter = {}
  for (const ch of p) {
    counter[ch] = (counter[ch] || 0) + 1
  }
  const result = []
  const plen = p.length
  const needs = Object.keys(counter).length
  let l = 0, r = 0
  const ctx = {}
  let valid = 0
  while (r < s.length) {
    const rch = s[r]
    r++
    if (!counter[rch]) continue
    ctx[rch] = (ctx[rch] || 0) + 1
    if (ctx[rch] !== counter[rch]) continue
    valid++
    while (valid === needs) {
      if (r - l === plen) {
        result.push(l)
      }
      const lch = s[l]
      if (counter[lch]) {
        if (ctx[lch] === counter[lch]) {
          valid--
        }
        ctx[lch] -= 1
      }
      l++
    }
  }
  return result
};