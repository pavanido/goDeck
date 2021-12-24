package apiserver

import (
	"net/http"
	"fmt"
	) 

var (
	num = 43
	sstr = "tik t0k"
)

func RequestHandler(res http.ResponseWriter, req *http.Request){
	fmt.Fprintf(res, "hello\n")
}

func StartServer(port string){
	http.ListenAndServe(port, nil)
}