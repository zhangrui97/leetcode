import "fmt"

func contains(s []uint64, e uint64) bool {
    for _, a := range s {
        if a == e {
            return true
        }
    }
    return false
}

func uniqueMorseRepresentations(words []string) int {
  morseMap := map[byte]uint64{
    'a': (2 << 56) | 0b01, //".-"
    'b': (4 << 56) | 0b1000, // "-..."
    'c': (4 << 56) | 0b1010, // "-.-."
    'd': (3 << 56) | 0b100, // "-.."
    'e': (1 << 56) | 0b0, // "."
    'f': (4 << 56) | 0b0010, // "..-."
    'g': (3 << 56) | 0b110, // "--."
    'h': (4 << 56) | 0b0000, // "...."
    'i': (2 << 56) | 0b00, // ".."
    'j': (4 << 56) | 0b0111, // ".---"
    'k': (3 << 56) | 0b101, // "-.-"
    'l': (4 << 56) | 0b0100, // ".-.."
    'm': (2 << 56) | 0b11, // "--"
    'n': (2 << 56) | 0b10, // "-."
    'o': (3 << 56) | 0b111, // "---"
    'p': (4 << 56) | 0b0110, // ".--."
    'q': (4 << 56) | 0b1101, // "--.-"
    'r': (3 << 56) | 0b010, // ".-."
    's': (3 << 56) | 0b000, // "..."
    't': (1 << 56) | 0b1, // "-"
    'u': (3 << 56) | 0b001, // "..-"
    'v': (4 << 56) | 0b0001, // "...-"
    'w': (3 << 56) | 0b011, // ".--"
    'x': (4 << 56) | 0b1001, // "-..-"
    'y': (4 << 56) | 0b1011, // "-.--"
    'z': (4 << 56) | 0b1100,  // "--.."
  }
  ans := make([]uint64, 0, len(words))
  for _, word := range words {
    num := uint64(0)
    bits := uint64(0)
    for i := range word {
      val := morseMap[word[i]]
      n := val >> 56
      num += n
      bits = (bits << n) | (val & 0xFF) 
    }
    result := (num << 56) | bits
    if !contains(ans, result) {
      ans = append(ans, result)
    }
  }
  return len(ans)
}