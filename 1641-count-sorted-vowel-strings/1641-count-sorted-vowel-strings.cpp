class Solution {
public:
  int countVowelStrings(int n) {
    int twoLetters = (n-1) * 10;
    int threeLetters = twoLetters*(n-2)/2;
    int fourLetters = threeLetters*(n-3)/6;
    int fiveLetters = fourLetters*(n-4)/20;
    return 5 + twoLetters + threeLetters + fourLetters + fiveLetters;
  }
};