/**
 * @param {string} s
 * @param {string} t
 * @return {boolean}
 */
var isSubsequence = function(s, t) {
  const isSub = (sl, sr, tl, tr) => {
    if (sl > sr) return true
    if (tl > tr) return false
    if (tl === tr) {
      return (sl === sr && t[tl] === s[sl])
    }
    if (sl === sr) {
      const ch = s[sl]
      for (; tl <= tr; tl++) {
        if (t[tl] === ch)
          return true
      }
      return false
    }
    for (; tl < tr; tl++) {
      if (t[tl] === s[sl])
        break
    }
    for (; tr > tl; tr--) {
      if (t[tr] === s[sr])
        break
    }
    return sl < sr && tl < tr && isSub(sl + 1, sr - 1, tl + 1, tr - 1)
  }
  return isSub(0, s.length - 1, 0, t.length - 1)
};