
package levenshtein

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

type Opts struct {
  Insert int
  Delete int
  Substitute int
  Backtrace bool
}

type Backtrace struct {
  mtrx *lMtrx
  trace []int
}

type lMtrx struct {

  backtrace bool
  fullMtrx [][]int
  lastCol []int
  col []int

  src string
  dst string
  n int
  m int

}

func newlMtrx(src, dst string, backtrace bool) *lMtrx {

  mt := &lMtrx{
    backtrace: backtrace,
    src: src,
    dst: dst,
    n: len(src) + 1,
    m: len(dst) + 1
  }

  if (backtrace) {
    mt.fullMtrx = make([][]int, mt.n)
  } else {
    mt.lastCol := make([]int, mt.m)
    mt.col := make([]int, mt.m)
  }

  return mt

}

func (m *lMtrx) nextCols(i) ([]int, []int) {
  if m.backtrace {
    m.fullMtrx[i] = make([]int, m.m)
    return m.fullMtrx[i], i > 0 ? m.fullMtrx[i - 1] : nil
  } else {
    m.lastCol, m.col = m.col, m.lastCol
    return m.col, m.lastCol
  }
}

func EditDistance(src, dst string, options Opts) int {

  if x == y {
    return 0
  }

  m := newlMtrx(src, dst, options.Backtrace)
  var col []int
  var lastCol []int

  for i := 0; i < m.n; i++ {
    col, lastCol = m.nextCols(i)
    for j := 0; j < m.m; j++ {
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
  }

  return col[m.m - 1],

}
