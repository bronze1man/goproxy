package main
/*
tcp代理
本程序用于想明白“双向”代理
A->C->B
A<-C<-B
做的事情是：
把A的数据读出来，写到B里面去。
tmp = a.read()
b.write(tmp)
把B的数据读出来，写到A里面去。
tmp = b.read()
a.write(tmp)
代理对应2个连接。每个对应2个方法，均实现即完成。
A的数据源是对象A自己
B的数据源是对象B自己
每个连接的数据源均为“对象”自己。对代理来说是对方。
xxxxxxxxxxxxxxx
*/
import (
"net"
"log"
"io"
"runtime"
"time"
//"bytes"
)
func dump(s string){
    //log.Print(s);
}
func proxy(c *net.TCPConn){
	defer c.Close();
	addr,err:=net.ResolveTCPAddr("tcp","localhost:8080");
	if (err!=nil){
		log.Print(err);
		return;
	}
	nextConn, err:= net.DialTCP("tcp",nil,addr)
	if (err!=nil){
		log.Print(err);
		return;
	}
	defer nextConn.Close();
	end :=make(chan int,1)
    //客户发送给服务器
	go func(){
		defer func(){end<-1}()
        defer nextConn.CloseWrite();
        defer c.CloseRead();
        defer dump("end half connect1 C");
		io.Copy(nextConn,c);
	}()
    //服务器发送给客户
	go func(){
		defer func(){end<-1}()
        defer nextConn.CloseRead();
        defer c.CloseWrite();
        defer dump("end half connect2 C");
		io.Copy(c,nextConn);
	}()
	<-end;
    <-end;
    dump("end connect C");
}

func listen(){
	addr,err:=net.ResolveTCPAddr("tcp",":8081");
	if (err!=nil){
		log.Print(err);
		return;
	}
	l, err := net.ListenTCP("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close();
	for {
		// Wait for a connection. 
		conn, err := l.AcceptTCP()
        dump("new connect C");
		if err != nil {
			log.Print(err)
			continue;
		}
		go proxy(conn)
	}
}
func server(){
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close();
	for {
		// Wait for a connection. 
		conn, err := l.Accept()
		if err != nil {
			log.Print(err)
			continue;
		}
        dump("new connect B");
		go func (c net.Conn){
            defer c.Close();
            defer dump("close connect B");
            readBuffer:=make([]byte,8*1024);
            for{
                _,err:=c.Read(readBuffer)
                if (err!=nil){
                    break;
                }
            }
            c.Write([]byte("1"));
        }(conn)
	}
}
func dumpGoRoutineNum(){
    for {
        time.Sleep(1000*1000*1000);
        log.Print(runtime.NumGoroutine());
        runtime.GC();
    }
}
func main(){
    go server();
    go dumpGoRoutineNum();
	listen();
}
