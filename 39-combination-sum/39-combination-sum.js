/**
 * @param {number[]} candidates
 * @param {number} target
 * @return {number[][]}
 */
var combinationSum = function(candidates, target) {
  const ans = []
  const backtrack = (cand, stack, sum) => {
    cand.forEach((c, i) => {
      if (c + sum < target) {
        backtrack(cand.slice(i), [...stack, c], sum + c)
      } else if (c + sum == target) {
        ans.push([...stack, c])  
      }
    })
  }
  backtrack(candidates, [], 0)
  return ans
};