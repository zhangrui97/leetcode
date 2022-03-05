class Solution:
  def removeDuplicates(self, nums: List[int]) -> int:
    l = len(nums)
    if l == 1: return 1
    slow = 0
    for i in range(1, l):
      if nums[i] != nums[slow]:
        slow += 1
        nums[slow] = nums[i]
    return slow + 1