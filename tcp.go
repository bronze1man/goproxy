package main

import "net"
import "time"
import "log"

//import "os"
import "io"
import "bytes"
import "fmt"
import "runtime"

var _ = io.Copy
var needEnd chan bool
func server() {
    l, err := net.Listen("tcp", ":20002")
    if err != nil {
        fmt.Print("1")
        log.Fatal(err)
    }
    defer l.Close()
    for {
        conn, err := l.Accept()
        if err != nil {
            log.Fatal(err)
        }
        conn.Close()
        go func(c net.Conn) {
            defer c.Close()
            io.Copy(c,c)
        }(conn)
    }
}
func client(b []byte) {
    defer func(){needEnd<-true}();
    time.Sleep(1*time.Second)
    conn, err := net.Dial("tcp", "127.0.0.1:20002")
    if err != nil {
        log.Fatal(err)
    }
    defer conn.Close()
    for i := 0; i < 10; i++ {
        conn.Write(b)
        
    }
}
func main() {
    fmt.Print("start\n")
    runtime.GOMAXPROCS(10)
    needEnd=make(chan bool)
    go server()
    time.Sleep(10 * time.Millisecond)
    totalThread := 10000;
    t0:=time.Now()
    for i:=0;i<totalThread;i++{
        go client(bytes.Repeat([]byte(`1`), 11))
    }
    t1:=time.Now()
    fmt.Printf("The call took %v to run.\n", t1.Sub(t0))
    for i:=0;i<totalThread;i++{
        <-needEnd
    }
    t2:=time.Now()
    fmt.Printf("The call took %v to run.\n", t2.Sub(t1))
    fmt.Print("finish\n")
}
