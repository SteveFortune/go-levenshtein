
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

// For each i = 1...m
//   For each j = 1...n
//                   { d(i-1,j) + 1
//     d(i,j) = min  { d(i,j-1) + 1
//                   { d(i-1,j-1) +  1; { if x(i) ≠ y(j)
//                                   0; { if x(i) = y(j)
func EditDistance(x, y string) int {

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

  costs := make([][]int, n)

  for i, _ := range costs {

    costs[i] = make([]int, m)

    for j, _ := range costs[i] {

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
        del := costs[lastI][j] + 1
        ins := costs[i][lastJ] + 1
        sub := costs[lastI][lastJ]
        if x[lastI] != y[lastJ] {
          sub++
        }
        cost = min([]int{del, ins, sub})
      }

      costs[i][j] = cost

    }
  }

  return costs[n - 1][m - 1]

}
