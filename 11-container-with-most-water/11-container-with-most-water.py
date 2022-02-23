class Solution:
    
    def maxArea(self, height: List[int]) -> int:
      def area(l, r) -> int:
        return (r-l)*min(height[l], height[r])
      l = 0
      r = len(height) - 1
      result = 0
      lastH = 0
      left = height[l]
      right = height[r]
      while (l < r):
        minH = min(left, right)
        if (minH > lastH):
          result = max(result, minH * (r - l))
          lastH = minH
        
        if left < right: 
          l += 1
          left = height[l]
        elif (left == right):
          l += 1
          left = height[l]
          r -= 1
          right = height[r]
        else: 
          r -= 1
          right = height[r]
        
      return result
        