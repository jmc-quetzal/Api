package main

import (
	"flag"
	"log"
	"net/http"
	"github.com/jmc-quetzal/api/routes"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Quetzal"))
}

func main() {
	r := routes.InitRouter()
	addr := flag.String("addr",":4000","HTTP Network address")
	flag.Parse()
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	log.Println("starting on :4000")
	err := http.ListenAndServe(*addr,r)
	log.Fatal(err)
}