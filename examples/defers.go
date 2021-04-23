package main

import (
	"fmt"
	"reflect"
	"time"
	//"database/sql/driver"

	"github.com/wwfyde/hello/gometry"
)

func main() {
	p := gometry.Point{0, 0}
	q := gometry.Point{3, 4}
	a := 1
	b := &a
	fmt.Println(a, b, *b, a == *b)
	distance := gometry.Distance(p, q)
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println(time.Now().String())

	fmt.Println(distance, reflect.TypeOf(distance))

}
