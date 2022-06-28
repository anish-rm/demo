package main
import "fmt"
func main(){
	fact:=1
	for i:=1;i<5;i++{
		fact=fact*i
}
	fmt.Println(fact)
}