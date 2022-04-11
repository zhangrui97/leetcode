int searchLeft(int* nums, int numsSize, int target, int l, int r){
  if (l == r) return l >= numsSize ? -1 : nums[l] == target ? l : -1;
  int m = (l + r) / 2;
  if (nums[m] < target) l = m + 1;
  else r = m;
  return searchLeft(nums, numsSize, target, l, r);
}
int searchRight(int* nums, int target, int l, int r){
  if (l == r) return r-1;
  int m = (l + r) / 2;
  if (nums[m] > target) r = m;
  else l = m + 1;
  return searchRight(nums, target, l, r);
}

/**
 * Note: The returned array must be malloced, assume caller calls free().
 */
int* searchRange(int* nums, int numsSize, int target, int* returnSize){
  int* result = (int*)malloc(2 * sizeof(int));
  *returnSize = 2;
  result[0] = -1;
  result[1] = -1;
  if (numsSize == 0) return result;
  int left = searchLeft(nums, numsSize, target, 0, numsSize);
  if (left != -1) {
    result[0] = left;
    result[1] = searchRight(nums, target, left+1, numsSize);
  }
  return result;
}