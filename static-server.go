package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

var port = flag.String("port", ":8090", "port for server to listen on")

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	log.Print(fmt.Sprintf("Listening on port :%s", *port))
	err := http.ListenAndServe(*port, nil)
	if err != nil {
		log.Fatal()
	}
}
