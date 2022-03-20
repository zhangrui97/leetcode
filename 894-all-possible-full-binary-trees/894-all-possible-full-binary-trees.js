/**
 * Definition for a binary tree node.
 * function TreeNode(val, left, right) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.left = (left===undefined ? null : left)
 *     this.right = (right===undefined ? null : right)
 * }
 */
/**
 * @param {number} n
 * @return {TreeNode[]}
 */
var allPossibleFBT = function(n) {
  if (n % 2 === 0) return []
  const level = (n + 1) / 2 
  const dp = new Array(level)
  dp[0] = [new TreeNode(0)]
  for (let i = 1; i < level; i ++) {
    dp[i] = []
    for (let l = 0; l < i; l ++) {
      dp[l].forEach(lNode => dp[i - 1 - l].forEach(rNode =>
        dp[i].push(new TreeNode(0, lNode, rNode))))
    }
  }
  return dp[level-1]
};