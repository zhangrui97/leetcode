/**
 * @param {string} allowed
 * @param {string[]} words
 * @return {number}
 */
var countConsistentStrings = function(allowed, words) {
  const set = new Set(allowed)
  const isConsistent = s => {
    for (let ch of s) {
      if (!set.has(ch)) return false
    }
    return true
  }
  let ans = 0
  for (word of words) {
    if (isConsistent(word)) ans++
  }
  return ans
};