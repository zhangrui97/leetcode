public class NumArray {
  int[] sums { get; }
  public NumArray(int[] nums) {
    sums = new int[nums.Length+1];
    int sum = 0;
    for (int i = 0; i < nums.Length; i++) {
      sum += nums[i];
      sums[i+1] = sum;
    }
  }

  public int SumRange(int left, int right) {
    return sums[right + 1] - sums[left];
  }
}

/**
 * Your NumArray object will be instantiated and called as such:
 * NumArray obj = new NumArray(nums);
 * int param_1 = obj.SumRange(left,right);
 */