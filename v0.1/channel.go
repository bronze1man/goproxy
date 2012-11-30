package main
//import "fmt"
import "time"
func main(){
    c:=make(chan int)
    for i:=0;i<210000;i++ {
        go func(i int){
        time.Sleep(1*time.Second)
        c<-i
        }(i)
    }
    //n:=0
    for i:=0;i<210000;i++{
        <-c
        //fmt.Print(n)
    }
}