package main
import "fmt"
func main(){
	var mark int
	fmt.Print("Enter mark : ")
	fmt.Scan(&mark)
	switch mark{
		case 100:fmt.Print("S")
		case 90:fmt.Print("A")
		case 80,85:fmt.Print("B")
		default:fmt.Print("Study well")
}
}