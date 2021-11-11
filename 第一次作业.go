package main

import (
	"fmt"
	"math/rand"
	"time"
)
func createnum(m []int){
	rand.Seed(time.Now().UnixNano())

	for i:=range m{
		m[i] = rand.Intn(1000)

	}

}
func bubblesort(n []int){

	for i:=0;i<len(n);i++ {

		for  j:=0;j<100-i-1;j++{
			if  temp:=0;n[j+1]<n[j]{
                 temp=n[j+1]
				 n[j+1]=n[j]
				 n[j]=temp

			}
		}
	}
}


func main01(){

i:=make([]int,100)
createnum(i)
fmt.Println("排序前：",i)
bubblesort(i)
fmt.Println("排序后：",i)

}
