object Solution {
  val table = Array(('I', 'V', 'X'), ('X', 'L', 'C'), ('C', 'D', 'M'), ('M', '?', '?'))
  def digitToRoman(digit: Int, position: Int): String = {
    val (one, five, ten) = table(position)
    digit match {
      case 0 => ""
      case 1 => s"$one"
      case 2 => s"$one$one"
      case 3 => s"$one$one$one"
      case 4 => s"$one$five"
      case 5 => s"$five"
      case 6 => s"$five$one"
      case 7 => s"$five$one$one"
      case 8 => s"$five$one$one$one"
      case 9 => s"$one$ten"
    }
  }
  def intToRoman(num: Int): String = {
    var result = ""
    var rest = num
    var pos = 0
    while (rest > 0) {
      result = s"${digitToRoman(rest%10, pos)}$result"
      pos = pos + 1
      rest = rest / 10
    }
    result
  }
}