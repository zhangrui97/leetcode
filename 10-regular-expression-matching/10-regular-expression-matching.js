/**
 * @param {string} s
 * @param {string} p
 * @return {boolean}
 */
var isMatch = function(s, p) {
  const cache = new Map()
  const sl = s.length
  const pl = p.length
  const dp = (si, pi) => {
    if (pi === pl) return si === sl;
    if (si === sl) {
      if ((pl - pi)%2 !== 0) return false
      for (let i = pi+1; i < pl; i += 2) {
        if (p[i] !== '*') return false
      }
      return true
    }
    
    const key = si * 1000 + pi
    if (cache.has(key)) return cache.get(key)
    
    let result = false
    if (s[si] === p[pi] || p[pi] === '.') {
      if (pi < pl - 1 && p[pi+1] === '*')
        result = dp(si, pi + 2) || dp(si + 1, pi) 
      else
        result = dp(si + 1, pi + 1)
    } else if (pi < pl - 1 && p[pi+1] === '*') {
      result = dp(si, pi + 2)
    }
    
    cache.set(key, result)
    return result
  }
  return dp(0, 0)
}