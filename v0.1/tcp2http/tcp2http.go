package main
/*
数据流图如下。
A->OurB->middleC->OurD->E
A<-OurB<-middleC<-OurD<-E

假设所有http请求都很普通，发送所有数据，处理，立刻返回所有数据。
    （不能长时间等待，不能将数据一部分一部分的返回。）
	目前已知非长连接效果很好。（一个tcp连接仅包含一个http请求，然后关闭tcp连接）

中间C仅转发http数据，不做其他事情。（一个标准http代理）
	
1.A->OurB(tcp)
2.A<-OurB(tcp)
    输入tcp请求
对象TCP，数据源A
	Serve 开启监听TCP。
	向数据源读Read读入数据
	向数据源写Write写入数据
	
3.OurB->middleC(http)
   B将数据，发送给中间人C
4.OurB<-middleC(http)
   B使用http推数据技术，轮询客户端，获取C发送的数据。
对象iHttp，数据源middleC或OurD
	向数据源读Read读入数据
	向数据源写Write写入数据
  
5.middleC->OurD(http)
   D获取到C的数据
6.middleC<-OurD(http)
   http轮询服务器端
对象oHTTP，数据源middleC或OurB
   
7.OurD->E
8.OurD<-E
还原tcp请求
对象TCP
	向数据源读Read读入数据
	向数据源写Write写入数据

*/
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


type iHTTP struct{
	writePoint int64   //下一次，写入的时候，写入的位置
	readPoint int64    //下一次，读取的时候，读取的位置
	canReadPoint int64 //当前一共可以读取的数据的位置
	writeSuccessPoint int64 //自己知道，对方成功读取的数据的位置
}
func (this *iHTTP) Write(p []byte)(n int,err error){
}
func readData(){
}
func listen(){
	// Listen on TCP port 2000 on all interfaces.
	l, err := net.Listen("tcp", ":2000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		// Wait for a connection. 
		conn, err := l.Accept()
		if err != nil {
			log.Print(err)
		}
		// Handle the connection in a new goroutine.
		// The loop then returns to accepting, so that
		// multiple connections may be served concurrently.
		go func(c net.Conn) {
			// Echo all incoming data.
			io.Copy(c, c)
			// Shut down the connection.
			c.Close()
		}(conn)
	}
}
func main(){
	net.Listent
}