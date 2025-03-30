package main

import "fmt"

type Counter struct {
	Value int
}

func modifyByValue(c Counter) {
	c.Value++
}

func modifyByPointer(c *Counter) {
	c.Value++
}

func main() {
	c1 := Counter{Value: 10}
	modifyByValue(c1)
	fmt.Println("After modifyByValue:", c1.Value) // ??
	// 10 is the answer .

	c2 := &Counter{Value: 10}
	modifyByPointer(c2)
	fmt.Println("After modifyByPointer:", c2.Value) // ??
	// 11 `is the answer..
}

// // ✅ Why?
// modifyByValue() receives a copy of Counter, so changes don’t affect the original.
// modifyByPointer() receives a pointer, so it modifies the original struct.









package main

import "fmt"

type Data struct {
	num int
}

func (d Data) nonPointerMethod() {
	d.num = 20
}

func (d *Data) pointerMethod() {
	d.num = 30
}

func main() {
	d := Data{num: 10}
	d.nonPointerMethod()
	fmt.Println("After nonPointerMethod:", d.num) // ??
	// 10 is the answer..



	d.pointerMethod()
	fmt.Println("After pointerMethod:", d.num) // ??
	// 30 is the answer..

	// ✅ Pointers: Why passing by reference is crucial for modifying struct data.
}
