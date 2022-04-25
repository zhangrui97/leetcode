public class Solution {
  public int SubarraySum(int[] nums, int k) {
    int counter = 0;
    var dict = new Dictionary<int, int>{[0] = 1};
    int sum = 0;
    foreach(int num in nums) {
      sum += num;
      if (dict.ContainsKey(sum - k)) {
        counter += dict[sum-k];
      }
      dict[sum] = dict.ContainsKey(sum) ? dict[sum] + 1 : 1;
    }
    
    return counter;
  }
}