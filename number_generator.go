package main

import (
    "fmt"
    "net/http"
    "strings"
    "log"
    "math/rand"
    "github.com/prometheus/client_golang/prometheus/promhttp"
    "strconv"
)

numberGen := rand.New(rand.NewSource(10000))

func randomNumber(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()
    for k, v := range r.Form {
        fmt.Println("key:", k)
        fmt.Println("val:", strings.Join(v, ""))
    }

    // // create seed
    // seed := int64(1000)
    max := int64(100)

    // parse args
    // if r.Form.Get("seed") != "" {
    // 	t, err := strconv.ParseInt(r.Form.Get("seed"), 10, 64)
    // 	if err != nil {
    // 		log.Fatal("Could not convert seed", err)
    // 		return
    // 	}
    // 	seed = t
    // }

    if r.Form.Get("max") != "" {
    	t, err := strconv.ParseInt(r.Form.Get("max"), 10, 64)
    	if err != nil {
    		log.Fatal("Could not convert max", err)
    		return
    	}
    	max = t
    }

    // generate number

    // send to client
    fmt.Fprintf(w, fmt.Sprint(numberGen.Int63n(max)))
}

func main() {
    http.HandleFunc("/", randomNumber) // set router
    http.Handle("/metrics", promhttp.Handler())
	log.Println("Serving on port 8080")
    err := http.ListenAndServe(":8080", nil) // set listen port
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}