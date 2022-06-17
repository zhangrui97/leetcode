/**
 * @param {number[]} nums
 * @return {number[][]}
 */
var subsets = function(nums) {
  const result = []
  const dfs = (path, i) => {
    result.push([...path])
    for (let j = i; j < nums.length; j++) {
      path.push(nums[j])
      dfs(path, j + 1)
      path.pop()
    }
  }
  dfs([], 0)
  return result  
};