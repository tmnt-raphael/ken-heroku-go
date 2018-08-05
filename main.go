package main

import (
  "log"
  "net/http"
  "os"
  "github.com/tmnt-raphael/kenwebutil"
)


func main() {
  port := os.Getenv("PORT")

  if port == "" {
    log.Fatal("$PORT must be set")
  }

  http.HandleFunc("/", kenwebutil.EchoPath)
  if err := http.ListenAndServe(":" + port, nil); err != nil {
    panic(err)
  }  

}