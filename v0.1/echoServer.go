package main
import(
    "net"
    "log"
    "io"
    "time"
    "runtime"
)
func server(){
	l, err := net.Listen("tcp", ":20002")
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
		go func (c net.Conn){
            defer c.Close();
            io.Copy(c,c);
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
    go dumpGoRoutineNum();
    server();
}