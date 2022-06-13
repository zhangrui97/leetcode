class Solution:
  def nextPermutation(self, nums: List[int]) -> None:
    """
    Do not return anything, modify nums in-place instead.
    """
    n = len(nums)
    k = n-2
    while k >= 0 and nums[k] >= nums[k+1]:
      k = k - 1
    if k < 0:
      nums[:] = nums[::-1]
    else:
      l = n-1
      v = nums[k]
      while l > k+1 and nums[l] <= v:
        l = l-1
      nums[k] = nums[l]
      nums[l] = v
      nums[k+1:] = nums[k+1:][::-1]
      
        