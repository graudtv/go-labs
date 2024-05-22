package main

import (
  "fmt"
  "sync"
)

/* set BUG to true to enable bug */
const BUG = false

type EntranceCounter struct {
  count int
  mx sync.Mutex
}

func (c *EntranceCounter) enter() {
  if (!BUG) {
    c.mx.Lock()
    defer c.mx.Unlock()
  }
  c.count++
}

func (c *EntranceCounter) getCount() int { return c.count }

func human(c *EntranceCounter) {
  for i := 0; i < 1000; i++ {
    c.enter()
  }
}

func main() {
  const humanCount = 100
  var c = EntranceCounter{}
  var wg sync.WaitGroup
  wg.Add(humanCount)

  for i := 0; i < humanCount; i++ {
    go func () {
      human(&c)
      wg.Done()
    }()
  }
  wg.Wait()
  fmt.Println("Total entrances:", c.getCount())
}
