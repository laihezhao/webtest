package main

import (
	"io"
	"net/http"
	"os/exec"
)

func reLanch()  {
	cmd := exec.Command("/bin/bash", "devops.sh")
	cmd.Start()
}

func hello(w http.ResponseWriter, r *http.Request)  {
	io.WriteString(w, "this is webtest")
	reLanch()
}

func main() {
	http.HandleFunc("/", hello)
	// 阿里云服务器上不能使用127.0.0.1
	err := http.ListenAndServe(":9001", nil)
	if err != nil {
		panic(err)
	}
}
