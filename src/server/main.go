package main

// Simple static server of Angular compiled dist/project folder.
import (
    "log"
    "net/http"
)

func main() {
    folder := "../client/dist-static/primeng-quickstart-cli/"
    http.Handle("/", http.FileServer(http.Dir(folder)))
    log.Fatal(http.ListenAndServe(":3000", nil))
}
