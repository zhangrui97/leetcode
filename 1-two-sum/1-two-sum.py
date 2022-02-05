class Solution:
  def twoSum(self, nums: List[int], target: int) -> List[int]:
    cache = {}
    for i in range(len(nums)):
      num = nums[i]
      if num in cache:
        return [cache[num], i]
      cache[target - num] = i