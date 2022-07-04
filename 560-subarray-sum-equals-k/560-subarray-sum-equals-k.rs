use std::collections::HashMap;

impl Solution {
  pub fn subarray_sum(nums: Vec<i32>, k: i32) -> i32 {
    let mut dict = HashMap::new();
    dict.insert(0, 1);
    let mut sum = 0;
    let mut counter = 0;
    for &v in nums.iter() {
      sum += v;
      let key = sum-k;
      if dict.contains_key(&key) {
        counter += dict[&key];
      }
      dict.insert(sum, if dict.contains_key(&sum) { dict[&sum]+1 } else { 1 });
    }
    counter
  }
}