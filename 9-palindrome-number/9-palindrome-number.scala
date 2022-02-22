object Solution {
  def isPalindrome(x: Int): Boolean = {
    if (x < 0 || (x%10==0 && x!=0)) return false
    var org = x
    var rev = 0
    while (org > rev) {
      rev = rev * 10 + org % 10
      org /= 10
    }
    return org == rev || rev /10 == org
  }
}