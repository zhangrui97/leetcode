struct NumMatrix {
  sum: Vec<Vec<i32>>
}


/** 
 * `&self` means the method takes an immutable reference.
 * If you need a mutable reference, change it to `&mut self` instead.
 */
impl NumMatrix {

  fn new(matrix: Vec<Vec<i32>>) -> Self {
    let m = matrix.len();
    let n = matrix[0].len();
    let mut sum: Vec<Vec<i32>> = Vec::new();
    sum.push(vec![0; n+1]);
    for (j, r) in matrix.iter().enumerate() {
      let pre = &sum[j];
      let mut row: Vec<i32> = Vec::new();
      row.push(0);
      for (i, &num) in r.iter().enumerate() {
        row.push(num+row[i]+pre[i+1]-pre[i]);
      }
      sum.push(row);
    }
    NumMatrix {
      sum
    }
  }

  fn sum_region(&self, row1: i32, col1: i32, row2: i32, col2: i32) -> i32 {
    let (r1, c1, r2, c2) = (row1 as usize, col1 as usize, (row2+1) as usize, (col2 + 1) as usize);
    self.sum[r2][c2] + self.sum[r1][c1] - self.sum[r2][c1] - self.sum[r1][c2]
  }
}

/**
 * Your NumMatrix object will be instantiated and called as such:
 * let obj = NumMatrix::new(matrix);
 * let ret_1: i32 = obj.sum_region(row1, col1, row2, col2);
 */