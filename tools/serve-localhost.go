package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
)

func main() {
	port := "8000"
	publicDir := "public"

	server := &http.Server{
		Addr:    ":" + port,
		Handler: http.FileServer(http.Dir(publicDir)),
	}

	go func() {
		fmt.Printf("Serving Go by Example at http://127.0.0.1:%s\n", port)
		server.ListenAndServe()
	}()

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if scanner.Text() == "exit" {
			fmt.Println("Shutting down...")
			server.Close()
			return
		}
	}
}
