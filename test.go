package main
import (
	"fmt"
	"strings"
)
func main(){
	var str string
	str="Hello world"
	fmt.Println(str)
	str1 := "Anish"
	fmt.Println(strings.ToUpper(str1))
	var a,b int = 5,10
	fmt.Println(a,b)
	c:=25.5
	fmt.Println("Value of a:",a)
	fmt.Printf("Type of c%T:\n",c)	
	fmt.Printf("Type of str1%T:\n",str1)	
	const i int = 25
	fmt.Println(i)	
	i:=30
	fmt.Println(i)	


}