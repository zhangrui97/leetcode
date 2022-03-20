object Solution {
  def countVowelStrings(n: Int): Int = {
    val (_, _, result) = ((5, 1, 0) /: (1 to 4)) {
      (s:Tuple3[Int, Int, Int], i:Int) => {
        val newCharCmb = s._1 * (5 - i) / (i + 1)
        val newPosCmb = s._2 * (n - i) / i
        (newCharCmb, newPosCmb, newCharCmb * newPosCmb + s._3)       
      } 
    }
    result + 5
  }
}