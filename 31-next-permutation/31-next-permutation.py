class Solution(object):
    def nextPermutation(self, nums):
        """
        :type nums: List[int]
        :rtype: None Do not return anything, modify nums in-place instead.
        """
        n = len(nums);
        k = n - 2
        while k >= 0 and nums[k] >= nums[k+1]:
          k = k - 1
        if k < 0:
          nums[:] = nums[::-1]
        else:
          v = nums[k]
          l = n - 1
          while l > k and nums[l] <= v:
            l = l - 1
          nums[l], nums[k] = v, nums[l]
          nums[k+1:] = nums[k+1:][::-1]

        