package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/sarathsp06/pasteit/internal/pasteit"
	"github.com/sarathsp06/pasteit/rpc/server"
)

func main() {
	addr := flag.String("addr", ":8081", "address to run the server")
	flag.Parse()
	twirpHandler := server.NewPasterServer(pasteit.NewServer(1, 1024*1024, nil), nil)
	if err := http.ListenAndServe(*addr, twirpHandler); err != nil {
		log.Println("error occurred", err)
	}
}
