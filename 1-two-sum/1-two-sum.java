import java.util.HashMap;

class Solution {
  public int[] twoSum(int[] nums, int target) {
    HashMap<Integer, Integer> cache = new HashMap<Integer, Integer>();
    for (int i = 0; i < nums.length; i++) {
      if (cache.containsKey(nums[i])) {
        return new int[] {cache.get(nums[i]), i};
      }
      cache.put(target - nums[i], i);
    }
    return new int[]{};
  }
}