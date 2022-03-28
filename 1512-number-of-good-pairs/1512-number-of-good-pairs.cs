public class Solution {
  public int NumIdenticalPairs(int[] nums) {
    var dict = new Dictionary<int, int>();
    int ans = 0;
    foreach (int n in nums) {
      if (dict.ContainsKey(n)) {
        ans += dict[n];
        dict[n]++;
      } else {
        dict[n] = 1;
      }
    }
    return ans;
  }
}