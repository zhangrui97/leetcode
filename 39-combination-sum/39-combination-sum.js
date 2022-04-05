/**
 * @param {number[]} candidates
 * @param {number} target
 * @return {number[][]}
 */
var combinationSum = function(candidates, target) {
  const ans = []
  const backtrack = (cand, stack, sum) => {
    for (let i = 0; i < cand.length; i++) {
      const v = cand[i]
      const next = sum + v
      if (next < target) {
        backtrack(cand.slice(i), [...stack, v], next)
      } else if (next === target) {
        stack.push(v)
        ans.push(stack)
      } else return
    }
  }
  candidates.sort((a, b) => a - b)
  backtrack(candidates, [], 0)
  return ans
};