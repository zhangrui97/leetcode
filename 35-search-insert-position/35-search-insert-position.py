class Solution:
  def searchInsert(self, nums: List[int], target: int) -> int:
    length = len(nums)
    def find(l, r):
      if nums[l] >= target: return l
      if nums[r] == target: return r
      if nums[r] < target: return r + 1
      mid = (l + r) // 2
      if nums[mid] < target: return find(mid+1, r)
      elif nums[mid] == target: return mid
      else: return find(l, mid-1)
    return find(0, length-1)