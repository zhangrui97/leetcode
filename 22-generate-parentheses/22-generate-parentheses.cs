public class Solution {
  List<string> result = new();
  private void BackTrack(string s, int left, int right) {
    if (left == 0 && right == 0) {
      result.Add(s);
    }
    if (left > 0)
      BackTrack(s+'(', left - 1, right);
    if (right > left)
      BackTrack(s+')', left, right - 1);
  }

  public IList<string> GenerateParenthesis(int n) {
    BackTrack("", n, n);
    return result;
  }
}