package main

import (
    "fmt"
    "net/http"
    "log"
    "math/rand"
    "github.com/prometheus/client_golang/prometheus/promhttp"
    "strconv"
    "time"
)


func randomNumber(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()
    max := int64(10000000000)
    rand.Seed(time.Now().UnixNano())
    
    if r.Form.Get("max") != "" {
    	t, err := strconv.ParseInt(r.Form.Get("max"), 10, 64)
    	if err != nil {
    		log.Println("Could not convert max", err.Error())
    		return
    	}
    	max = t
    }
    // send to client
    fmt.Fprintf(w, fmt.Sprint(rand.Int63n(max)))
}

func main() {
    http.HandleFunc("/randomNumber", randomNumber) // set router
    http.Handle("/metrics", promhttp.Handler())
	log.Println("Serving on port 8080")
    err := http.ListenAndServe(":8080", nil) // set listen port
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}