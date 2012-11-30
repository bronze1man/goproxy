package main
import(
"net/http/httputil"
"net/http
)
/*
2012/08/01 00:33:43 http: StatusNotModified response with header "Content-Type" defined
return 304...
*/
func noopDirector(req *http.Request){
    log.Print(req.URL)
}
func main(){
    handler := &httputil.ReverseProxy{};
    handler.Director=noopDirector
    rw := &http.ResponseWriter{}
    
    handler.ServeHTTP(
}