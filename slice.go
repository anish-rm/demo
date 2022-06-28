package main
import "fmt"
func main(){
	m1:=make([]int,5)
	m2:=[] int {1,2,3}
	fmt.Println(m2[2])
	fmt.Printf("M1=%v\n",m1)
	fmt.Printf("length=%d\n",len(m1))
	fmt.Printf("Cap=%d\n",cap(m1))
}