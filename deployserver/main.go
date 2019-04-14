package main

import (
	"io"
	"log"
	"net/http"
	"os/exec"
)

func reLaunch() {
	cmd := exec.Command("sh", "./deploy.sh")
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	err = cmd.Wait()
}

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "<h1>Hello World! this is deploy server</h1>")
	reLaunch()
}

func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe(":7000", nil)
}
