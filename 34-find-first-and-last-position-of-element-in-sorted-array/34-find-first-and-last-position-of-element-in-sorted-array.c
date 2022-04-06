int search(int* nums, int target, int* l, int* r){
  if (*l > *r) return -1;
  int m = (*l + *r) / 2;
  if (nums[m] == target) return m;
  if (nums[m] > target) *r = m - 1;
  else *l = m + 1;
  return search(nums, target, l, r);
}
int searchLeft(int* nums, int target, int l, int r){
  if (l == r) return l;
  int m = (l + r) / 2;
  if (nums[m] < target) l = m + 1;
  else r = m;
  return searchLeft(nums, target, l, r);
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
  int l = 0, r = numsSize - 1;
  int mid = search(nums, target, &l, &r);
  if (mid == -1) {
    result[0] = -1;
    result[1] = -1;
  } else {
    result[0] = searchLeft(nums, target, l, mid);
    result[1] = searchRight(nums, target, mid, r + 1);
  }
  return result;
}