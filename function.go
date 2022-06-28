package main
import "fmt"
import "time"
func main(){
go myfun("First")
time.Sleep(time.Second)
go myfun("Second")
go myfun("third")
time.Sleep(time.Microsecond)
fmt.Println("Complted")
}
func myfun(a string){
for i:=0;i<3;i++{
fmt.Println(a," : ",i)
}
}
//create multiple goroutines which prints name of the author, id of the author