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
  if (n === 1) return [new TreeNode(0)]
  const ans = []
  for (let l = 1; l < n; l += 2) {
    allPossibleFBT(l).forEach(lNode =>
      allPossibleFBT(n-l-1).forEach(rNode =>
        ans.push(new TreeNode(0, lNode, rNode))))
  }
  return ans
};