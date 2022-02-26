class Solution:
  def letterCombinations(self, digits: str) -> List[str]:
    dic = {
      "2": "abc",
      "3": "def",
      "4": "ghi",
      "5": "jkl",
      "6": "mno",
      "7": "pqrs",
      "8": "tuv",
      "9": "wxyz"
    }
    def combine(aoa: List[Iterable[str]]) -> List[str]:
      if not aoa: return []
      if len(aoa) == 1: return aoa[0]
      head, *tail = aoa
      return [h+t for h in head for t in combine(tail)]
    return combine(list(map(lambda d: dic[d], digits)))