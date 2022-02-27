class Solution:
    def isValid(self, s: str) -> bool:
      if len(s) % 2 > 0: return False
      stack = []
      for ch in s:
        if ch == '(': stack += ')'
        elif ch == '[': stack += ']'
        elif ch == '{': stack += '}'
        else:
          if len(stack) == 0 or stack.pop() != ch:
            return False
        
      return len(stack) == 0
    
        