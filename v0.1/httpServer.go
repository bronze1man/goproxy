package main

import(
"net/http"
"log"
"io"
)
var _=io.Copy
func t1(){
    //b:=[]byte(`Hello world`)
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	    //w.Write(b)
	    io.WriteString(w,`Hello world`)
    })
}
func t2(){
    b:=[]byte(`Hello world`)
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	    w.Write(b)
    })
}
func main(){
    t2()

    log.Fatal(http.ListenAndServe(":8080", nil))
}
