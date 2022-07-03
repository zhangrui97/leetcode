impl Solution {
  pub fn search(nums: Vec<i32>, target: i32) -> i32 {
    let mut lo: usize = 0;
    let mut hi: usize = nums.len() - 1;
    while lo <= hi {
      let mid = (lo + hi) / 2;
      if nums[mid] == target { return mid as i32 }
      if nums[mid] > target {
        if mid == 0 { 
          return -1 
        } else {
          hi = mid - 1;
        }
      } else {
        lo = mid + 1;
      }
    }
    -1
  }
}