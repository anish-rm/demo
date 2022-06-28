package main
import "fmt"
func main(){
	var a = [5] int {1,2,3}
	fmt.Println(a)
	b:=[3] int {4,5,6}
	fmt.Println(b)
	c:=[5] int {1:50,4:100}
	fmt.Println(c)
	a[3]=5
	fmt.Println(a)
}
