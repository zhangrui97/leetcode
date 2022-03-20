impl Solution {
  pub fn count_vowel_strings(n: i32) -> i32 {
    let x : i32 = n * (n + 5);
    return (x * (x + 10) + 24) / 24;
  }
}