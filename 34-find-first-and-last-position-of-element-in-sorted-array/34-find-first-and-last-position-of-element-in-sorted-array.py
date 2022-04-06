class Solution:
  def searchRange(self, nums: List[int], target: int) -> List[int]:
    def search() -> Tuple[int, int, int]:
      l, r = 0, len(nums)-1
      while l <= r:
        mid = (l+r)//2
        if nums[mid] == target: return l, mid, r
        if nums[mid] > target: 
          r = mid - 1
        else:
          l = mid + 1
      return 0, -1, 0
    def searchLeft(l, r) -> int:
      while l < r:
        mid = (l+r)//2
        if nums[mid] < target:
          l = mid + 1
        else:
          r = mid
      return l
    def searchRight(l, r) -> int:
      while l < r:
        mid = (l+r)//2
        if nums[mid] > target:
          r = mid
        else:
          l = mid + 1
      return l - 1
    l, m, r = search()
    if m == -1: return [-1, -1]
    else: return [searchLeft(l, m), searchRight(m, r+1)]