/**
 * Definition for a binary tree node.
 * function TreeNode(val, left, right) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.left = (left===undefined ? null : left)
 *     this.right = (right===undefined ? null : right)
 * }
 */
/**
 * @param {TreeNode} root
 * @return {number[]}
 */
var inorderTraversal = function(root) {
  const ans = []
  const dfs = t => {
    if (t == null) { return }
    dfs(t.left)
    ans.push(t.val)
    dfs(t.right)
  } 
  dfs(root)
  return ans
};