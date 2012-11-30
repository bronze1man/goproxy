package main
import(
"net/http"
"time"
"log"
"io"
//"net/http/httputil"
)

var proxyTransport http.RoundTripper = &http.Transport{}
type httpProxyHandler struct{
	Transport http.RoundTripper
};

func copyHeader(dst, src http.Header) {
	for k, vv := range src {
		for _, v := range vv {
			dst.Add(k, v)
		}
	}
}
func (p *httpProxyHandler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	transport := p.Transport
	if transport == nil {
		transport = http.DefaultTransport
	}

	outreq := new(http.Request)
	*outreq = *req // includes shallow copies of maps, but okay

	outreq.Proto = "HTTP/1.1"
	outreq.ProtoMajor = 1
	outreq.ProtoMinor = 1
	outreq.Close = true

	outreq.Header = make(http.Header)
	copyHeader(outreq.Header, req.Header)
	if outreq.Header.Get("Connection") != "" {
		outreq.Header.Del("Connection")
	}

	res, err := transport.RoundTrip(outreq)
	if err != nil {
		log.Printf("http: proxy error: %v", err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	
	//write response back
	
	copyHeader(rw.Header(), res.Header)
	rw.WriteHeader(res.StatusCode)

	if res.Body != nil {
		var dst io.Writer = rw
		io.Copy(dst, res.Body)
	}
}

func proxy(){
    handler := &httpProxyHandler{};
    s := &http.Server{
        Addr:           ":8080",
        Handler:        handler,
        ReadTimeout:    100 * time.Second,
        WriteTimeout:   100 * time.Second,
        MaxHeaderBytes: 1 << 20,
    }
    log.Fatal(s.ListenAndServe());
}
func main(){
	net.DialTCP
    proxy();
}