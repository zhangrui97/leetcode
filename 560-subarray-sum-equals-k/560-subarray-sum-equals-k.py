class Solution:
  def subarraySum(self, nums: List[int], k: int) -> int:
    map = {0 : 1}
    sum = 0
    ans = 0
    for v in nums:
      sum += v
      ans += map.get(sum-k, 0)
      map[sum] = map.get(sum, 0) + 1
    return ans
      
    