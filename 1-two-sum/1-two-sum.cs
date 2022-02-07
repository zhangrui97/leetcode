public class Solution {
    public int[] TwoSum(int[] nums, int target) {
        var cache = new Dictionary<int,int>();
        for (int i = 0; i < nums.Length; i++) {
          if (cache.ContainsKey(nums[i])) {
            return new int[]{cache[nums[i]], i};
          }
          cache[target - nums[i]] = i;
        }
        return null;
    }
}