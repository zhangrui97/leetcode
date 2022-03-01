class Solution {
public:
  bool isValid(string s) {
    if (s.size() % 2) return false;
    std::map<char, char> pair = {{'(',')'},{'[',']'},{'{','}'}};
    std::stack<char> stack;
    for (char ch : s) {
      if (pair.count(ch) > 0) stack.push(pair[ch]);
      else {
        if (stack.empty() || stack.top() != ch) return false;
        else stack.pop();
      }
    }
    return stack.empty();
  }
};