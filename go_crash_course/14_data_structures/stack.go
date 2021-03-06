package main

import (
  "fmt"
  "container/list"
)

func main() {
  customQueue := &customStack{
    stack: list.New(),
  }

  fmt.Printf("Push: A\n")
  customQueue.Push("A")
  fmt.Printf("Push: B\n")
  customQueue.Push("B")
  fmt.Printf("Size: %d\n", customQueue.Size())
  for customQueue.Size() > 0 {
    frontVal, _ := customQueue.Front()
    fmt.Printf("Front: %s\n", frontVal)
    fmt.Printf("Pop: %s\n", frontVal)
    customQueue.Pop()
  }
  fmt.Printf("Size: %d\n", customQueue.Size())
}

type customStack struct {
  stack *list.List
}

func (c *customStack) Push(value string) {
  c.stack.PushFront(value)
}

func (c *customStack) Pop() error {
  if c.stack.Len() > 0 {
    el := c.stack.Front()
    c.stack.Remove(el)
  }
  return fmt.Errorf("Pop Error: Queue is empty")
}

func (c *customStack) Front() (string, error) {
  if c.stack.Len() > 0 {
    if val, ok := c.stack.Front().Value(string); ok {
      return val, nil
    }

    return "", fmt.Errorf("Peep Error: Queue datatype is incorrect")
  }

  return "", fmt.Errorf("Peep Error: Queue is empty")
}

func (c *customStack) Size() int {
  return c.stack.Len()
}

func (c *customStack) Empty() {
  return c.stack.Len() == 0
}
