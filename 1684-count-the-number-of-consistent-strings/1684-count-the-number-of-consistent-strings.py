from functools import reduce
class Solution:
  def countConsistentStrings(self, allowed: str, words: List[str]) -> int:
    return reduce(lambda acc, word: acc+1 if all(ch in allowed for ch in word) else acc, words, 0)
        