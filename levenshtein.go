
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
  InsWeight int
  DelWeight int
  SubWeight int
  Backtrace bool
}

var DefaultOpts = Opts{ 1, 1, 1, false }

type Backtrace struct {
  mtrx *lMtrx
  trace []int
}

type lMtrx struct {

  backtrace bool

  fullMtrx [][]int
  lastCol []int
  col []int

  n int
  m int

}

func newlMtrx(n, m int, backtrace bool) *lMtrx {

  mt := &lMtrx{
    backtrace: backtrace,
    n: n + 1,
    m: m + 1,
  }

  if (backtrace) {
    mt.fullMtrx = make([][]int, mt.n)
  } else {
    mt.lastCol = make([]int, mt.m)
    mt.col = make([]int, mt.m)
  }

  return mt

}

func (m *lMtrx) nextCol(i int) {

  if m.backtrace {

    m.col = make([]int, m.m)
    m.fullMtrx[i] = m.col
    if i > 0 {
      m.lastCol = m.fullMtrx[i - 1]
    }

  } else {
    m.lastCol, m.col = m.col, m.lastCol
  }

}

func EditDistance(src, dst string, options Opts) (int, *Backtrace) {

  if src == dst {
    return 0, nil
  }

  var n = len(src)
  var m = len(dst)

  if n == 0 {
    return m, nil
  }
  if m == 0 {
    return n, nil
  }

  mt := newlMtrx(n, m, options.Backtrace)

  for i := 0; i < mt.n; i++ {
    mt.nextCol(i)
    for j := 0; j < mt.m; j++ {
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
        del := mt.lastCol[j] + options.DelWeight
        ins := mt.col[lastJ] + options.InsWeight
        sub := mt.lastCol[lastJ]
        if src[lastI] != dst[lastJ] {
          sub += options.SubWeight
        }
        cost = min([]int{del, ins, sub})
      }
      mt.col[j] = cost
    }
  }

  return mt.col[mt.m - 1], &Backtrace{}

}
