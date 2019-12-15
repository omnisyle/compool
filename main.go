package main

type Object interface {
  Close()
}

type Factory func() (Object, error)

type ObjectPool struct {
  capacity int64
  objects  chan Object
  factory  Factory
}

func NewObjectPool(factory Factory, capacity int) *ObjectPool {
  return &ObjectPool{
    capacity: int64(capacity),
    factory:  factory,
    objects:  make(chan Object, capacity),
  }
}
