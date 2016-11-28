package main

import (
    "net/http"
    "encoding/json"

    "github.com/vetbase/husky"
)

type Message struct {
    Text string
}

func index(w http.ResponseWriter, r *http.Request) {
    m := Message{"Welcome to Husky!"}
    b, err := json.Marshal(m)

    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.Write(b)
}

func main() {
    service := husky.New()
    service.GET("/", index)
    service.Serve("8080")
}
