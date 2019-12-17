package main

import "errors"

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

func (p *ObjectPool) Get() (Object, error) {
	obj, ok := <-p.objects
	if !ok {
		return nil, errors.New("Pool is closed")
	}

	return obj, nil
}

func (p *ObjectPool) Put(obj Object) error {
	p.objects <- obj
	return nil
}
