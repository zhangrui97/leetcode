class Solution {
  private int searchRec(int[] nums, int target, int lo, int hi) {
    if (lo > hi) return -1;
    int mid = (lo + hi) / 2;
    int v = nums[mid];
    if (v == target) { return mid; }
    if (v > target) { return searchRec(nums, target, lo, mid-1); }
    return searchRec(nums, target, mid+1, hi);
  }
  public int search(int[] nums, int target) {
    return searchRec(nums, target, 0, nums.length-1);
  }
}