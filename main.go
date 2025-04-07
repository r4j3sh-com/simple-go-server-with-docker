package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	newServeMux := http.NewServeMux()

	newServeMux.HandleFunc("/", handlePage)

	const port = ":8080"
	srv := http.Server{
		Handler:      newServeMux,
		Addr:         port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Println("Server started on", port)
	err := srv.ListenAndServe()
	log.Fatal(err)
}

func handlePage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	const page = `<!DOCTYPE html>
	<head></head>
	<body>
	    <h1>Welcome to the Simple Web Server!</h1>
		<p>This is a simple web server written in Go.</p>
		<p>Feel free to navigate around.</p>
		<h2>Hello from Docker! I'm a Go server.</h2>
	</body>
	</html>`

	w.Write([]byte(page))
}
