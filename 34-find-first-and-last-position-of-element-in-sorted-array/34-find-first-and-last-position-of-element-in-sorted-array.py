class Solution:
  def searchRange(self, nums: List[int], target: int) -> List[int]:
    def search(l, r) -> Tuple[int, int, int]:
      if l > r: return 0, -1, 0
      mid = (l+r)//2
      if nums[mid] == target: return l, mid, r
      if nums[mid] > target: return search(l, mid-1)
      else: return search(mid + 1, r)
    def searchLeft(l, r) -> int:
      if l == r: return l
      mid = (l+r)//2
      if nums[mid] < target: return searchLeft(mid + 1, r)
      else:return searchLeft(l, mid)
    def searchRight(l, r) -> int:
      if l == r: return l-1
      mid = (l+r)//2
      if nums[mid] > target: return searchRight(l, mid)
      else: return searchRight(mid + 1, r)
    l, m, r = search(0, len(nums)-1)
    if m == -1: return [-1, -1]
    else: return [searchLeft(l, m), searchRight(m, r+1)]