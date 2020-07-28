package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
        "strconv"
        "time"
)

const ( version = "v1.0.0" )

func humanize(n int64) string {
    in := strconv.FormatInt(n, 10)
    numOfDigits := len(in)
    if n < 0 {
        numOfDigits-- // First character is the - sign (not a digit)
    }
    numOfCommas := (numOfDigits - 1) / 3

    out := make([]byte, len(in)+numOfCommas)
    if n < 0 {
        in, out[0] = in[1:], '-'
    }

    for i, j, k := len(in)-1, len(out)-1, 0; ; i, j = i-1, j-1 {
        out[j] = in[i]
        if i == 0 {
            return string(out)
        }
        if k++; k == 3 {
            j, k = j-1, 0
            out[j] = ','
        }
    }
}

func headers(w http.ResponseWriter, req *http.Request) {
    log.Println("Requested Header API")
    for name, headers := range req.Header {
        for _, h := range headers {
            fmt.Fprintf(w, "%v: %v\n", name, h)
            log.Println(fmt.Sprintf("%v: %v", name, h))
        }
    }
}


func main() {

	log.Println("Starting Timer Test...")

        http.HandleFunc("/version", func(w http.ResponseWriter, r *http.Request) { fmt.Fprintf(w, version)
                                                                                   log.Println("Requested Version API")
                                                                                 })
        http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) { w.WriteHeader(200)
                                                                                  log.Println("Requested Health API")
                                                                                })
        http.HandleFunc("/header", headers )

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

            pp := r.URL.Path[1:len(r.URL.Path)]
            k, err := strconv.Atoi(pp) 
            k = k * 1000000000
            if ( (len(pp) < 1) || (err != nil) ) { 
                fmt.Fprint(w,"Try http://127.0.0.1:3000/5 to have me count up to 5 billion and tell you how long it took.") 
                return
              } else {
            
            then := time.Now()            
            for i := 1; i<=k; i++ { }
            now := time.Now()
            diff := now.Sub(then);
            s_diff := diff.Seconds();
            str_diff := fmt.Sprintf("%f",s_diff)
            fmt.Fprint(w,"I counted to " + humanize(int64(k)) + " in " + str_diff + " seconds") 
            log.Println("I counted to " + humanize(int64(k)) + " in  " + str_diff + " seconds")
              }
	})

	var port string
	port = os.Getenv("PORT")
        if ( len(port) == 0 ) {
		port = "3000"
	}

        log.Println("Using port"+port)

	log.Fatal(http.ListenAndServe(":"+port, nil))

	s := http.Server{Addr: ":" + port }

	go func() { log.Fatal(s.ListenAndServe()) }()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	<-signalChan

	log.Println("Shutdown signal received, exiting...")

	s.Shutdown(context.Background())
}
