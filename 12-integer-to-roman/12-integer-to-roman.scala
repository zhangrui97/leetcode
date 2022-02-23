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
  val d2r = collection.immutable.HashMap(
    (0, 0) -> digitToRoman(0, 0),
    (0, 1) -> digitToRoman(1, 0),
    (0, 2) -> digitToRoman(2, 0),
    (0, 3) -> digitToRoman(3, 0),
    (0, 4) -> digitToRoman(4, 0),
    (0, 5) -> digitToRoman(5, 0),
    (0, 6) -> digitToRoman(6, 0),
    (0, 7) -> digitToRoman(7, 0),
    (0, 8) -> digitToRoman(8, 0),
    (0, 9) -> digitToRoman(9, 0),
    (1, 0) -> digitToRoman(0, 1),
    (1, 1) -> digitToRoman(1, 1),
    (1, 2) -> digitToRoman(2, 1),
    (1, 3) -> digitToRoman(3, 1),
    (1, 4) -> digitToRoman(4, 1),
    (1, 5) -> digitToRoman(5, 1),
    (1, 6) -> digitToRoman(6, 1),
    (1, 7) -> digitToRoman(7, 1),
    (1, 8) -> digitToRoman(8, 1),
    (1, 9) -> digitToRoman(9, 1),
    (2, 0) -> digitToRoman(0, 2),
    (2, 1) -> digitToRoman(1, 2),
    (2, 2) -> digitToRoman(2, 2),
    (2, 3) -> digitToRoman(3, 2),
    (2, 4) -> digitToRoman(4, 2),
    (2, 5) -> digitToRoman(5, 2),
    (2, 6) -> digitToRoman(6, 2),
    (2, 7) -> digitToRoman(7, 2),
    (2, 8) -> digitToRoman(8, 2),
    (2, 9) -> digitToRoman(9, 2),
    (3, 0) -> digitToRoman(0, 3),
    (3, 1) -> digitToRoman(1, 3),
    (3, 2) -> digitToRoman(2, 3),
    (3, 3) -> digitToRoman(3, 3),
  )
  def intToRoman(num: Int): String = {
    var result = ""
    var rest = num
    var pos = 0
    while (rest > 0) {
      result = s"${d2r(pos, rest%10)}$result"
      pos = pos + 1
      rest = rest / 10
    }
    result
  }
}