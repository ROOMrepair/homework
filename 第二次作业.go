package main

import "fmt"

type str struct {
	a int
	b string
	c bool
}
func Receiver(v interface{}){

	switch v.(type) {
	case int:
		fmt.Printf("%d是int",v)
	case string:
		fmt.Printf("%s是string",v)
	case bool:
		fmt.Printf("%v是bool",v)
	case str:
		fmt.Printf("%v是struct",v)
	}
}
func main02(){
	var i =str{7,"hj",true}

	Receiver(i)


}