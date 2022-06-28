package main
import "fmt"
func main(){
	var a int
	fmt.Println("enter a : ")
	fmt.Scan(&a)
	fmt.Print("Value is\n",a)
	if(a%2==0){
		fmt.Print("It is even")
	}else{
		fmt.Print("It is odd")
	}
}