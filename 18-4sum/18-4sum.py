class Solution:
  def fourSum(self, nums: List[int], target: int) -> List[List[int]]:
    l = len(nums)
    if l < 4: return []
    if l == 4 and sum(nums) == target: return [nums]
    nums.sort()
    avg = target / 4
    if nums[0] >= avg + 1: return []
    if nums[-1] < avg: return []
    
    result = []
    a = 0
    while a < l-3:
      na = nums[a]
      b = a + 1
      while b < l-2:
        nb = nums[b]
        rest = target - na - nb
        c = b + 1
        d = l - 1
        while c < d:
          nc = nums[c]
          nd = nums[d]
          s = nc + nd
          if s < rest:
            while c < d and nums[c] == nc: c += 1
          elif s == rest:
            result += [[na, nb, nc, nd]]
            while c < d and nums[c] == nc: c += 1
            while c < d and nums[d] == nd: d -= 1
          else:
            while c < d and nums[d] == nd: d -= 1
        while b < l-2 and nums[b] == nb: b += 1
      while a < l-3 and nums[a] == na: a += 1
    return result