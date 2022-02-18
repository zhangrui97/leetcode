class Solution:
  def threeSum(self, nums: List[int]) -> List[List[int]]:
    ln = len(nums)
    if ln < 3: return []
    nums.sort()
    if nums[0] > 0 or nums[-1] < 0: return []
    if nums[0] == 0 and nums[-1] == 0: return [[0, 0, 0]]
    result = []
    i = 0
    while i < ln-2:
      a = nums[i]
      if a > 0: return result
      l = i + 1
      r = ln - 1
      while l < r:
        b = nums[l]
        c = nums[r]
        if c < 0: return result
        if a + b + c < 0:
          l += 1
          while l < r and nums[l] == b: l += 1
        elif a + b + c == 0:
          result.append([a, b, c])
          l += 1
          r -= 1
          while l < r and nums[l] == b: l += 1
          while l < r and nums[r] == c: r -= 1
        else:
          r -= 1
          while l < r and nums[r] == c: r -= 1
      i += 1
      while i < ln-2 and nums[i] == a: i += 1

    return result