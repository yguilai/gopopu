package gopopu

import (
	"fmt"
	"testing"
	"time"
)

type Test struct {
	IntData   int
	FloatData float32
	S         *Test3
}

type Test2 struct {
	IntData   int
	FloatData float32
	Text      string
	S         *Test3
}

type Test3 struct {
	Data int
}

func TestPopulate(t *testing.T) {
	var aa = &Test{}
	var bb = &Test2{IntData: 11, FloatData: 22, Text: "test", S: &Test3{Data: 22}}
	t1 := time.Now()
	err := Populate(aa, bb)
	if err != nil {
		panic(err)
	}
	t2 := time.Now()
	fmt.Println(aa)
	fmt.Println(t2.Sub(t1))
}
