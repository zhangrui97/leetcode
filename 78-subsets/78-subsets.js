/**
 * @param {number[]} nums
 * @return {number[][]}
 */
var subsets = function(nums) {
  const result = []
  const dfs = (path, i) => {
    const res = []
    for (let j = 0, p = path; p > 0; j++, p = p >> 1) {
      if (p & 1) res.push(nums[j])
    }
    result.push(res)
    for (let j = i; j < nums.length; j++) {
      dfs(path | (1 << j), j + 1)
    }
  }
  dfs(0, 0)
  return result  
};