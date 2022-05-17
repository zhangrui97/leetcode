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
 * @param {number} key
 * @return {TreeNode}
 */
var deleteNode = function(root, key) {
  if (root == null) return null
  if (root.val === key) {
    if (!root.left && !root.right) return null
    if (root.left && root.right) {
      let min = root.right
      while (min.left) { min = min.left }
      root.right = deleteNode(root.right, min.val)
      min.left = root.left
      min.right = root.right
      return min
    }
    if (root.left) return root.left
    if (root.right) return root.right
  } 
  if (key < root.val) root.left = deleteNode(root.left, key)
  else root.right = deleteNode(root.right, key)
  return root
};