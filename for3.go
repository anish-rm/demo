package main
import "fmt"
func main(){
	number := [3]int{10,20,30}
	for i,x:= range number{
		fmt.Printf("x %d is at %d \n",x,i)
}
}