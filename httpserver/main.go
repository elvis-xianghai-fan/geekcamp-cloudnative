package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

const noContent = 204
const ok = 200

func main() {
	log.Println("This is a test HTTP server.")
	os.Setenv("VERSION", "1.0")

	http.HandleFunc("/", defaultHandler)
	http.HandleFunc("/healthz", healthz)

	err := http.ListenAndServe(getListeningAddr(), nil)
	if err != nil {
		log.Fatal(err)
	}
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Client IP: %s", getIP(r))

	for k, vs := range r.Header {
		log.Printf("%s = %s", k, vs)
		for _, v := range vs {
			w.Header().Add(k, v)
		}
	}

	w.Header().Set("VERSION", os.Getenv("VERSION"))
	w.WriteHeader(noContent)

	log.Printf("Responded with: %d", noContent)
}

func healthz(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "ok")
	w.WriteHeader(ok)
}

func getIP(r *http.Request) string {
	forwarded := r.Header.Get("X-FORWARDED-FOR")
	if forwarded != "" {
		return forwarded
	}
	return r.RemoteAddr
}

func getListeningAddr() string {
	servicePort := os.Getenv("SERVICE_PORT")
	if len(servicePort) == 0 {
		servicePort = "80"
	}
	return fmt.Sprintf(":%s", servicePort)
}
