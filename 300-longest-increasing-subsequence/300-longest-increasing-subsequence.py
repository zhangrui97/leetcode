class Solution:
  def lengthOfLIS(self, nums: List[int]) -> int:
    buffer = []
    for n in nums:
      l = 0
      r = len(buffer)
      while l < r:
        m = (l + r) // 2
        if buffer[m] < n: l = m + 1
        else: r = m
      if l >= len(buffer):
        buffer.append(n)
      else:
        buffer[l] = n
    return len(buffer)