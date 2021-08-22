package main

import (
	"flag"

	_ "github.com/neimv/backend-go-leo/color"
)

var (
	host = flag.String("h", "localhost", "Host to start web server")
	port = flag.Int("p", 3000, "port to access web server")
)

func main() {
	color.Colorize("Starting server")
}
