package main

import (
	"flag"
	"fmt"

	server "github.com/neimv/backend-go-leo/server"
	utils "github.com/neimv/backend-go-leo/utilities"
	views "github.com/neimv/backend-go-leo/views"
)

var (
	host = flag.String("h", "localhost", "Host to start web server")
	port = flag.Int("p", 3000, "port to access web server")
)

func main() {
	url := fmt.Sprintf("%s:%d", *host, *port)
	fmt.Println("Hello")
	utils.PrintBlue("Starting WebServer")

	server := server.NewServer(url)
	server.Handle("GET", "/", views.HandleRootRequest)
	server.RunServer()
}
