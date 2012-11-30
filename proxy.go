package main
import (
"net"
"log"
"runtime/debug"
//"io"
"bytes"
"net/http"
"bufio"
"fmt"
)
func tIsPanic(err error){
    if (err==nil){
        return;
    }
    debug.PrintStack()
    log.Panic(err)
}
func tIsError(err error)bool{
    if (err==nil){
        return false;
    }
    fmt.Println(err)
    return true;
}
func tRecoverPanic(){
                if x := recover(); x != nil {
                    log.Printf("run time panic: %v", x)
                }
}
//
func loadReader(r io.Reader)([]byte,io.Reader){
    bbuf:=bytes.NewBuffer(nil)
    io.Copy(bbuf,r)
    data:=bbuf.Bytes()
    
}
func server(){
    l,err:=net.Listen("tcp",":20001")
    tIsPanic(err);
    defer l.Close();
    for{
        conn,err:=l.Accept();
        tIsPanic(err);
        go func(c net.Conn) {
            defer func(){
                c.Close()
            }()

            
            buf:=bufio.NewReader(c)
            req,err:=http.ReadRequest(buf)
            if tIsError(err){
                return;
            }
            fmt.Println(req.Method,req.URL.String())
            transport:=&http.Transport{};
            resp,err:=transport.RoundTrip(req)
            if tIsError(err){
                return;
            }
            err=resp.Write(c)
            if tIsError(err){
                return;
            }
        }(conn)
    }
}
func main(){
    fmt.Print("start")
    server()
}