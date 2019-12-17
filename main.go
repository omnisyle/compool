package main

import (
	"fmt"
	"math/rand"
	"time"
)

type FakeObj struct {
	id int
}

func (f *FakeObj) Close() {}

func factory() (Object, error) {
	src := rand.NewSource(time.Now().Unix())
	ran := rand.New(src)
	return &FakeObj{
		id: ran.Int(),
	}, nil
}

func main() {
	objPool := NewObjectPool(factory, 5)
	obj, _ := factory()
	err := objPool.Put(obj)

	if err != nil {
		panic(err)
	}

	obj, err = objPool.Get()

	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", obj)

	obj, err = objPool.Get()

	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", obj)
}
