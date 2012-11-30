package main
import "net"
import "time"
import "log"
import "os"
import "io"
import "fmt"
func server(){
    addr,err:=net.ResolveUDPAddr("udp",":20001");
    if (err!=nil){
        log.Fatal(err);
    }
    conn,err:=net.ListenUDP("udp",addr)
    if (err!= nil){
        fmt.Print("1")
        log.Fatal(err)
    }
    defer conn.Close();
    for{
        io.Copy(os.Stdout,conn);
    }
}
func client(b []byte){
    conn,err:=net.Dial("udp","127.0.0.1:20001")
    if (err!=nil){
        log.Fatal(err)
    }
    defer conn.Close();
    for{
        conn.Write(b);
    }
}
func main(){
    go server();
    time.Sleep(1*time.Second);
    go client([]byte(`1`));
    go client([]byte(`2`));
    time.Sleep(1*time.Second);
}