impl Solution {
  pub fn search(nums: Vec<i32>, target: i32) -> i32 {
    let mut ans = -1;
    let (mut lo, mut hi) = (0, nums.len());
    while lo < hi {
      let mid = (lo + hi) / 2;
      if nums[mid] == target { 
        ans = mid as i32;
        break;
      }
      if nums[mid] > target {
        hi = mid;
      } else {
        lo = mid + 1;
      }
    }
    ans
  }
}