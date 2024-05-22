package main

import (
  "fmt"
)

type TemperatureRange struct {
  min int
  max int
}

var InitialRange = TemperatureRange{min: 15, max: 30}

func (tr *TemperatureRange) IsValid() bool {
  return tr.min <= tr.max
}

func (tr *TemperatureRange) AdjustUpperLimit(limit int) {
  if limit < tr.max {
    tr.max = limit
  }
}

func (tr *TemperatureRange) AdjustLowerLimit(limit int) {
  if limit > tr.min {
    tr.min = limit
  }
}

func main() {
  var N, K, lim int
  var op string

  fmt.Scan(&N)
  for n := 0; n < N; n++ {
    var tr = InitialRange
    fmt.Scan(&K)
    for k := 0; k < K; k++ {
      fmt.Scan(&op)
      fmt.Scan(&lim)
      if op == ">=" {
        tr.AdjustLowerLimit(lim)
      } else if op == "<=" {
        tr.AdjustUpperLimit(lim)
      } else {
        panic("Bad input")
      }

      if (tr.IsValid()) {
        fmt.Println(tr.min)
      } else {
        fmt.Println(-1)
      }
    }
  }
}
