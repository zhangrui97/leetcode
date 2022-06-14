

void nextPermutation(int* nums, int numsSize){
  int k = numsSize - 2;
  for (; k >= 0; k--) {
    if (nums[k+1]>nums[k]) break;
  }
  if (k >= 0) {
    int l = numsSize - 1;
    int v = nums[k];
    for (; l > k+1; l--) {
      if (nums[l]>v) break;
    }
    nums[k] = nums[l];
    nums[l] = v;
  }
  for (int i = k+1, j = numsSize-1; i < j; i++, j--) {
    int temp = nums[i];
    nums[i] = nums[j];
    nums[j] = temp;
  }
}