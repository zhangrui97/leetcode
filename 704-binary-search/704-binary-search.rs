use std::cmp::Ordering;

fn search(nums: &Vec<i32>, target: i32, lo: usize, hi: usize) -> i32 {
  if lo > hi { 
    -1
  } else {
    let mid = (lo+hi) / 2;
    match nums[mid].cmp(&target) {
      Ordering::Less => search(nums, target, mid+1, hi),
      Ordering::Greater => if mid == 0 { -1 } else { search(nums, target, lo, mid-1) },
      Ordering::Equal => mid as i32,
    }
  }
}

impl Solution {
  pub fn search(nums: Vec<i32>, target: i32) -> i32 {
    search(&nums, target, 0, nums.len()-1)
  }
}