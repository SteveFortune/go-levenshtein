
package levelshtein

func min(arr []int) int {
  if len(arr) == 0 {
    return 0
  }
  min := arr[0]
  for _, v := range arr {
    if v < min {
      min = v
    }
  }
  return min
}

// Calls `EditDistanceCal` with the default substitution
// calibration
//
func EditDistance(x, y string) int {
  return EditDistanceWeight(x, y, 0)
}

// For each i = 1...m
//   For each j = 1...n
//                   { d(i-1,j) + 1
//     d(i,j) = min  { d(i,j-1) + 1
//                   { d(i-1,j-1) +  1; { if x(i) ≠ y(j)
//                                   0; { if x(i) = y(j)
func EditDistanceWeight(x, y string, subWeight int) int {

  if x == y {
    return 0
  }

  var n = len(x) + 1
  var m = len(y) + 1

  if m == 0 {
    return n
  }
  if n == 0 {
    return m
  }

  if subWeight == 0 {
    subWeight = 1
  }

  lastCol := make([]int, m)
  col := make([]int, m)

  for i := 0; i < n; i++ {
    for j := 0; j < m; j++ {
      var cost int
      if i == 0 && j == 0 {
        cost = 0
      } else if i == 0 {
        cost = j
      } else if j == 0 {
        cost = i
      } else {
        lastI := i - 1
        lastJ := j - 1
        del := lastCol[j] + 1
        ins := col[lastJ] + 1
        sub := lastCol[lastJ]
        if x[lastI] != y[lastJ] {
          sub += subWeight
        }
        cost = min([]int{del, ins, sub})
      }
      col[j] = cost
    }
    lastCol, col = col, lastCol
  }

  return lastCol[m - 1]

}
