package main
import "net"
import "time"
import "log"
//import "os"
import "io"
import "bytes"
import "fmt"
var _=io.Copy
func server(){
    l,err:=net.Listen("tcp",":20001")
    if (err!= nil){
        fmt.Print("1")
        log.Fatal(err)
    }
    defer l.Close();
    for{
        conn,err := l.Accept();
        if err!=nil{
            log.Fatal(err)
        }
        go func(c net.Conn){
            defer c.Close();
            fmt.Print(`+`);
            //var b bytes.Buffer;
            b:=bytes.NewBuffer([]byte{});
            io.Copy(b,c);
            data:=b.Bytes();
            fmt.Print(len(data));
            fmt.Print(`-`);
        }(conn)
    }
}
func client(b []byte){
    conn,err:=net.Dial("tcp","127.0.0.1:20001")
    if (err!=nil){
        log.Fatal(err)
    }
    defer conn.Close();
    for i:=0;i<1000;i++{
        conn.Write(b);
            time.Sleep(1*time.Millisecond);
        }
    fmt.Print(`finish`)

}
func main(){
    go server();
    time.Sleep(1*time.Second);
    go client(bytes.Repeat([]byte(`1`),1000));
    //go client([]byte(`2`));
    time.Sleep(1*time.Second);
}