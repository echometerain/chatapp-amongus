package main

import (
	"os"
	"os/signal" //good practice for handling stop signals
	"syscall"
)

var sig = make(chan os.Signal, 1) //shuts down the program nicely
func main() { // init function
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	go discord()
}
