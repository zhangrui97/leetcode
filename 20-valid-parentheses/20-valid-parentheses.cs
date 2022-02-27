public class Solution {
  private static Dictionary<char, char> map = new () {
    {'{','}'}, {'[',']'}, {'(',')'}
  };
  public bool IsValid(string s) {
    Stack<char> stack = new();
    foreach (char ch in s) {
      if (map.ContainsKey(ch)) {
        stack.Push(map[ch]);
      }
      else {
        if (stack.Count == 0 || stack.Pop() != ch)
          return false;
      }
    }
    return stack.Count == 0;
  }
}