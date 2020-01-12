package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

var (
	allowRequestPerSecond int
	hostname              string
	port                  int
)

func init() {
	flag.IntVar(&allowRequestPerSecond, "limit", 10, "allow request per second")
	flag.IntVar(&port, "port", 3000, "listen port")
	flag.StringVar(&hostname, "hostname", "localhost", "ip address or hostname")

	flag.CommandLine.Usage = func() {
		o := flag.CommandLine.Output()
		fmt.Fprintf(o, "Usage: %s\n", flag.CommandLine.Name())
		fmt.Fprintf(o, "A simple ratelimiter\n")
		flag.PrintDefaults()
	}

	flag.Parse()

	// Initialize RateLimiter middleware
	InitRateLimiter(allowRequestPerSecond)
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func main() {
	log.Println("allowed number of requests: ", allowRequestPerSecond)

	addr := fmt.Sprintf("%s:%v", hostname, port)

	router := http.NewServeMux()
	router.HandleFunc("/", handler)
	if err := http.ListenAndServe(addr, RateLimitMiddleware(router)); err != nil {
		log.Panicf("[FAILED] an error has occured at starting server: %v", err)
	}
}
