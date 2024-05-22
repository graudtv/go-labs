package main

import (
  "fmt"
  "container/heap"
)

type IntHeap []int

/* Implementation of heap.Interface:
 * - sort.Interface:
 *   - Len() int
 *   - Less(i, j int) bool
 *   - Swap(i, j int)
 * - Push(x any)
 * - Pop() any
 */
func (h IntHeap) Len() int { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x any) { *h = append(*h, x.(int)) }
func (h *IntHeap) Pop() any {
  res := (*h)[len(*h) - 1]
  *h = (*h)[:len(*h) - 1]
  return res
}

func main() {
  var N, x, k, res int
  h := &IntHeap{}
  fmt.Scan(&N)
  for i := 0; i < N; i++ {
    fmt.Scan(&x)
    h.Push(x)
  }
  fmt.Scan(&k)

  heap.Init(h)
  for i := 0; i < k; i ++ {
    res = heap.Pop(h).(int)
  }
  fmt.Println(res)
}
