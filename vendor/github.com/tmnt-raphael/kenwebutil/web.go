package kenwebutil

import (
  "net/http"
)

func EchoPath(w http.ResponseWriter, r *http.Request) {
  message := r.URL.Path

  w.Write([]byte(message))
}