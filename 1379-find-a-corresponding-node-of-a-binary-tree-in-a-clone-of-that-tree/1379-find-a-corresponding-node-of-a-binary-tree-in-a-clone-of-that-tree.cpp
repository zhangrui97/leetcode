/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     TreeNode *left;
 *     TreeNode *right;
 *     TreeNode(int x) : val(x), left(NULL), right(NULL) {}
 * };
 */

class Solution {
public:
  TreeNode* getTargetCopy(TreeNode* original, TreeNode* cloned, TreeNode* target) {
    if (original == NULL || original == target) return cloned;
    TreeNode* result = getTargetCopy(original->left, cloned->left, target);
    if (result != NULL) {
      return result;
    } else {
      return getTargetCopy(original->right, cloned->right, target);
    }
  }
};