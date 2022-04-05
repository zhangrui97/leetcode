/**
 * @param {number[]} candidates
 * @param {number} target
 * @return {number[][]}
 */
var combinationSum = function(candidates, target) {
  const ans = []
  const backtrack = (cand, stack, sum) => {
    for (let i = 0; i < cand.length; i++) {
      const s = [...stack, cand[i]]
      const next = sum + cand[i]
      if (next < target) {
        backtrack(cand.slice(i), s, next)
      } else if (next === target) {
        ans.push(s)
      } else return
    }
  }
  candidates.sort((a, b) => a - b)
  backtrack(candidates, [], 0)
  return ans
};